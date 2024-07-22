package dao

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/labring/sealos/service/account/helper"

	"github.com/labring/sealos/controllers/pkg/types"
)

func TestCockroach_GetPayment(t *testing.T) {
	db, err := NewAccountInterface("", "", "")
	if err != nil {
		t.Fatalf("NewAccountInterface() error = %v", err)
		return
	}
	got, err := db.GetPayment(types.UserQueryOpts{Owner: "1fgtm0mn"}, time.Time{}, time.Time{})
	if err != nil {
		t.Fatalf("GetPayment() error = %v", err)
		return
	}
	t.Logf("got = %+v", got)
}

func TestMongoDB_GetAppCosts(t *testing.T) {
	db, err := NewAccountInterface("", "", "")
	if err != nil {
		t.Fatalf("NewAccountInterface() error = %v", err)
		return
	}
	appCosts, err := db.GetAppCosts(&helper.AppCostsReq{
		UserBaseReq: helper.UserBaseReq{
			TimeRange: helper.TimeRange{
				StartTime: time.Now().Add(-24 * time.Hour * 30),
				EndTime:   time.Now(),
			},
			Auth: helper.Auth{
				Owner: "xxx",
			},
		},
		Namespace: "ns-xxx",
		AppType:   "APP",
		AppName:   "xxx",
		Page:      72,
		PageSize:  10,
	})
	if err != nil {
		t.Fatalf("GetAppCosts() error = %v", err)
		return
	}
	t.Logf("appCosts = %+v", appCosts)
}

func TestCockroach_GetTransfer(t *testing.T) {
	os.Setenv("LOCAL_REGION", "97925cb0-c8e2-4d52-8b39-d8bf0cbb414a")

	db, err := NewAccountInterface("", "", "")
	if err != nil {
		t.Fatalf("NewAccountInterface() error = %v", err)
		return
	}
	transfer, err := db.GetTransfer(&types.GetTransfersReq{
		UserQueryOpts: &types.UserQueryOpts{
			Owner: "q0xeg9z1",
		},
		Type: 0,
		LimitReq: types.LimitReq{
			Page:     1,
			PageSize: 10,
			TimeRange: types.TimeRange{
				StartTime: time.Now().UTC().Add(-1*time.Hour - 30*time.Minute),
				EndTime:   time.Now().UTC(),
			},
		},
	})
	if err != nil {
		t.Fatalf("GetTransfer() error = %v", err)
		return
	}
	t.Logf("timerange = %+v", types.TimeRange{
		StartTime: time.Now().UTC().Add(-30*time.Hour - 30*time.Minute),
		EndTime:   time.Now().UTC(),
	})
	t.Logf("transfer = %+v", transfer.LimitResp)
}

func TestMongoDB_GetCostAppList(t *testing.T) {
	dbCTX := context.Background()
	m, err := NewAccountInterface("", "", "")
	if err != nil {
		t.Fatalf("NewAccountInterface() error = %v", err)
		return
	}
	defer func() {
		if err = m.Disconnect(dbCTX); err != nil {
			t.Errorf("failed to disconnect mongo: error = %v", err)
		}
	}()
	req := helper.GetCostAppListReq{
		Auth: helper.Auth{
			Owner: "hwhbg4vf",
		},
		//Namespace: "ns-hwhbg4vf",
		AppType: "APP-STORE",
		//AppName: "cronicle-ldokpaus",
		LimitReq: helper.LimitReq{
			Page:     1,
			PageSize: 5,
		},
	}
	appList, err := m.GetCostAppList(req)
	if err != nil {
		t.Fatalf("failed to get cost app list: %v", err)
	}
	t.Logf("len costAppList: %v", len(appList.Apps))
	t.Logf("costAppList: %#+v", appList)
}

func TestMongoDB_GetCostOverview(t *testing.T) {
	dbCTX := context.Background()
	m, err := NewAccountInterface("", "", "")
	if err != nil {
		t.Fatalf("NewAccountInterface() error = %v", err)
		return
	}
	defer func() {
		if err = m.Disconnect(dbCTX); err != nil {
			t.Errorf("failed to disconnect mongo: error = %v", err)
		}
	}()
	req := helper.GetCostAppListReq{
		Auth: helper.Auth{
			Owner: "hwhbg4vf",
		},
		//Namespace: "ns-hwhbg4vf",
		AppType: "APP-STORE",
		//AppName: "cronicle-ldokpaus",
		LimitReq: helper.LimitReq{
			Page:     1,
			PageSize: 5,
		},
	}
	appList, err := m.GetCostOverview(req)
	if err != nil {
		t.Fatalf("failed to get cost app list: %v", err)
	}
	t.Logf("len costAppList: %v", len(appList.Overviews))
	t.Logf("costAppList: %#+v", appList)
}
