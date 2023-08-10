#!/bin/bash
set -ex

cloudDomain="127.0.0.1.nip.io"
tlsCrtPlaceholder="<tls-crt-placeholder>"
tlsKeyPlaceholder="<tls-key-placeholder>"
mongodbUri=""
saltKey=""

function prepare {
  # kubectl apply namespace, secret and mongodb
  kubectl apply -f manifests/namespace.yaml

  # apply notifications crd
  kubectl apply -f manifests/notifications_crd.yaml

  # gen mongodb uri
  gen_mongodbUri

  # gen saltKey if not set or not found in secret
  gen_saltKey

  # mutate desktop config
  mutate_desktop_config

  # create tls secret
  create_tls_secret
}

function gen_mongodbUri() {
  # if mongodbUri is empty then create mongodb and gen mongodb uri
  if [ -z "$mongodbUri" ]; then
    echo "no mongodb uri found, create mongodb and gen mongodb uri"
    kubectl apply -f manifests/mongodb.yaml
    # if there is no sealos-mongodb-conn-credential secret then wait for mongodb ready
    while [ -z "$(kubectl get secret -n sealos sealos-mongodb-conn-credential)" ]; do
      echo "waiting for mongodb secret generated"
      sleep 5
    done
    chmod +x scripts/gen-mongodb-uri.sh
    mongodbUri=$(scripts/gen-mongodb-uri.sh)
  fi
}

function gen_saltKey() {
    password_salt=$(kubectl get secret desktop-frontend-secret -n sealos -o jsonpath="{.data.password_salt}" 2>/dev/null || true)
    if [[ -z "$password_salt" ]]; then
        saltKey=$(tr -dc 'a-z0-9' </dev/urandom | head -c64 | base64 -w 0)
    else
        saltKey=$password_salt
    fi
}

function mutate_desktop_config() {
  if kubectl get secret desktop-frontend-secret -n sealos --ignore-not-found > /dev/null 2>&1; then
    echo "desktop-frontend-secret already exists, skip mutate desktop secret"
  else
    # mutate etc/sealos/desktop-config.yaml by using mongodb uri and two random base64 string
    sed -i -e "s;<your-mongodb-uri-base64>;$(echo -n "$mongodbUri" | base64 -w 0);" etc/sealos/desktop-config.yaml
    sed -i -e "s;<your-jwt-secret-base64>;$(tr -cd 'a-z0-9' </dev/urandom | head -c64 | base64 -w 0);" etc/sealos/desktop-config.yaml
    sed -i -e "s;<your-password-salt-base64>;$saltKey;" etc/sealos/desktop-config.yaml
  fi
}

function create_tls_secret {
  if grep -q $tlsCrtPlaceholder manifests/tls-secret.yaml; then
    echo "mock tls secret"
    kubectl apply -f manifests/mock-cert.yaml
    echo "mock tls cert has been created successfully."
  else
    echo "tls secret is already set"
    kubectl apply -f manifests/tls-secret.yaml
  fi
}

function sealos_run_controller {
  # run user controller
  sealos run tars/user.tar \
  --env cloudDomain=$cloudDomain \
  --env cloudPort="6443"

  # run terminal controller
  sealos run tars/terminal.tar \
  --env cloudDomain=$cloudDomain \
  --env userNamespace="user-system" \
  --env wildcardCertSecretName="wildcard-cert" \
  --env wildcardCertSecretNamespace="sealos-system"

  # run app controller
  sealos run tars/app.tar

  # run resources monitoring controller
  sealos run tars/monitoring.tar \
  --env MONGO_URI="$mongodbUri" --env DEFAULT_NAMESPACE="resources-system"

  # run resources metering controller
  sealos run tars/metering.tar \
  --env MONGO_URI="$mongodbUri" --env DEFAULT_NAMESPACE="resources-system"

  # run account controller
  sealos run tars/account.tar \
  --env MONGO_URI="$mongodbUri" \
  --env DEFAULT_NAMESPACE="account-system"

  # run licenseissuer controller
  sealos run tars/licenseissuer.tar \
  --env canConnectToExternalNetwork="true" \
  --env enableMonitor="true"
}

function sealos_authorize {
    echo "start to authorize sealos"
    echo "create admin-user"
    # create admin-user
    kubectl apply -f manifests/admin-user.yaml
    # wait for admin-user ready
    echo "waiting for admin-user generated"
    while true; do
        if kubectl get namespace ns-admin >/dev/null 2>&1 && kubectl get accounts.account.sealos.io admin -n sealos-system >/dev/null 2>&1; then
            break
        else
            echo "waiting for preset admin-user to be created..."
            sleep 3
        fi
    done
    # issue license for admin-user
    echo "license issue for admin-user"
    kubectl apply -f manifests/free-license.yaml
}


function sealos_run_frontend {
  echo "run desktop frontend"
  configFileFlag=""
  if kubectl get secret desktop-frontend-secret -n sealos --ignore-not-found > /dev/null 2>&1; then
    configFileFlag=""
  else
    configFileFlag="--config-file etc/sealos/desktop-config.yaml"
  fi
  sealos run tars/frontend-desktop.tar \
    --env cloudDomain=$cloudDomain \
    --env certSecretName="wildcard-cert" \
    --env passwordEnabled="true" \
    $configFileFlag

  echo "run applaunchpad frontend"
  sealos run tars/frontend-applaunchpad.tar \
  --env cloudDomain=$cloudDomain \
  --env certSecretName="wildcard-cert"

  echo "run terminal frontend"
  sealos run tars/frontend-terminal.tar \
  --env cloudDomain=$cloudDomain \
  --env certSecretName="wildcard-cert"

  echo "run dbprovider frontend"
  sealos run tars/frontend-dbprovider.tar \
  --env cloudDomain=$cloudDomain \
  --env certSecretName="wildcard-cert"

  echo "run cost center frontend"
  sealos run tars/frontend-costcenter.tar \
  --env cloudDomain=$cloudDomain \
  --env certSecretName="wildcard-cert" \
  --env transferEnabled="true" \
  --env rechargeEnabled="false"
}

function install {
  # gen mongodb uri and others
  prepare

  # sealos run controllers
  sealos_run_controller
  
  # sealos run frontends
  sealos_run_frontend

  # sealos authorize
  sealos_authorize
}

install
