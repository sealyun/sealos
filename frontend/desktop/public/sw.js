if(!self.define){let e,s={};const c=(c,i)=>(c=new URL(c+".js",i).href,s[c]||new Promise((s=>{if("document"in self){const e=document.createElement("script");e.src=c,e.onload=s,document.head.appendChild(e)}else e=c,importScripts(c),s()})).then((()=>{let e=s[c];if(!e)throw new Error(`Module ${c} didn’t register its module`);return e})));self.define=(i,a)=>{const n=e||("document"in self?document.currentScript.src:"")||location.href;if(s[n])return;let r={};const t=e=>c(e,n),o={module:{uri:n},exports:r,require:t};s[n]=Promise.all(i.map((e=>o[e]||t(e)))).then((e=>(a(...e),r)))}}define(["./workbox-3a642b46"],(function(e){"use strict";importScripts(),self.skipWaiting(),e.clientsClaim(),e.precacheAndRoute([{url:"/_next/static/Pk0sl4J8KwseDp6Zd1uwV/_buildManifest.js",revision:"65b64feb16ffdbafce8ea39a06cdc3b7"},{url:"/_next/static/Pk0sl4J8KwseDp6Zd1uwV/_ssgManifest.js",revision:"b6652df95db52feb4daf4eca35380933"},{url:"/_next/static/chunks/012ff928-5fd5ae6c3e73c42c.js",revision:"5fd5ae6c3e73c42c"},{url:"/_next/static/chunks/10e5c763-6ed24905adf99715.js",revision:"6ed24905adf99715"},{url:"/_next/static/chunks/161-311b22822b646f15.js",revision:"311b22822b646f15"},{url:"/_next/static/chunks/17-5f73c18f2592ace1.js",revision:"5f73c18f2592ace1"},{url:"/_next/static/chunks/264.692582dfc1508e18.js",revision:"692582dfc1508e18"},{url:"/_next/static/chunks/324-64222386cc30314e.js",revision:"64222386cc30314e"},{url:"/_next/static/chunks/345-ae90393d6fbf56d9.js",revision:"ae90393d6fbf56d9"},{url:"/_next/static/chunks/485-8c8e2b28d7b79d73.js",revision:"8c8e2b28d7b79d73"},{url:"/_next/static/chunks/658.d578ee96127a9af8.js",revision:"d578ee96127a9af8"},{url:"/_next/static/chunks/668-877224645c132965.js",revision:"877224645c132965"},{url:"/_next/static/chunks/693.0793abdc13fea790.js",revision:"0793abdc13fea790"},{url:"/_next/static/chunks/926.fde13c3972bfd87f.js",revision:"fde13c3972bfd87f"},{url:"/_next/static/chunks/framework-ea81dd6b0e8bf8e4.js",revision:"ea81dd6b0e8bf8e4"},{url:"/_next/static/chunks/main-afc5f8c9e927184b.js",revision:"afc5f8c9e927184b"},{url:"/_next/static/chunks/pages/404-132225303e165308.js",revision:"132225303e165308"},{url:"/_next/static/chunks/pages/WorkspaceInvite-716681175fc103f0.js",revision:"716681175fc103f0"},{url:"/_next/static/chunks/pages/_app-534b4b9ab81d69fc.js",revision:"534b4b9ab81d69fc"},{url:"/_next/static/chunks/pages/_error-1307886b7255ca63.js",revision:"1307886b7255ca63"},{url:"/_next/static/chunks/pages/callback-72b538675f32686c.js",revision:"72b538675f32686c"},{url:"/_next/static/chunks/pages/index-f6005ed6f7b95db4.js",revision:"f6005ed6f7b95db4"},{url:"/_next/static/chunks/pages/proxyOAuth-fc155941f5c64ef7.js",revision:"fc155941f5c64ef7"},{url:"/_next/static/chunks/pages/signin-1d6520ee380e2424.js",revision:"1d6520ee380e2424"},{url:"/_next/static/chunks/pages/switchRegion-2a9a2903f1c157f7.js",revision:"2a9a2903f1c157f7"},{url:"/_next/static/chunks/polyfills-c67a75d1b6f99dc8.js",revision:"837c0df77fd5009c9e46d446188ecfd0"},{url:"/_next/static/chunks/webpack-5c3e6b319c6caf2c.js",revision:"5c3e6b319c6caf2c"},{url:"/_next/static/css/0236986349b1075d.css",revision:"0236986349b1075d"},{url:"/_next/static/css/492a2d4bba196b6c.css",revision:"492a2d4bba196b6c"},{url:"/_next/static/css/8fb6cedfc06155f6.css",revision:"8fb6cedfc06155f6"},{url:"/_next/static/css/9c5ab77e948405b5.css",revision:"9c5ab77e948405b5"},{url:"/_next/static/css/f29d1b5c7c2cc422.css",revision:"f29d1b5c7c2cc422"},{url:"/_next/static/media/close_white.5b1bded4.svg",revision:"c5b8c262b90daa29e85c67e95ed4b6d7"},{url:"/_next/static/media/lock.92d8e0c3.svg",revision:"ca1cde5a3cc61da7ee9bfb8cc7e824fe"},{url:"/_next/static/media/warning.feae7a43.svg",revision:"3e0578143b73e97deee3bcc5140da84b"},{url:"/favicon.ico",revision:"e0adb2be6bc609982b5161e42609a99f"},{url:"/iconfont/iconfont.js",revision:"9d01c1074cd56c3bda4409e8fdfa0dc6"},{url:"/icons/apps.svg",revision:"563e729e5619ccf33b4331812eb275f1"},{url:"/icons/close.png",revision:"d5d067253444824c88a22a920bebaa5f"},{url:"/icons/close_white.svg",revision:"c5b8c262b90daa29e85c67e95ed4b6d7"},{url:"/icons/driverStar.svg",revision:"d4cc375236ce7a5257937e83be0480b6"},{url:"/icons/empty.svg",revision:"3d60593509734249bde21dc710fd3746"},{url:"/icons/favicon-16x16.png",revision:"9fba9b4339260ac3a3bd90581e34e7e2"},{url:"/icons/favicon-32x32.png",revision:"a7253d6b434315a9a007a611a0396791"},{url:"/icons/home.svg",revision:"c1538af0548f7bff4213c754862d33b0"},{url:"/icons/icon-512x512.png",revision:"8a3598942cc5a7d14f0296b083cf92c0"},{url:"/icons/license.svg",revision:"82aa5984c926eca5dcb9e90ef206a987"},{url:"/icons/maximize.png",revision:"a92404a02a73ee9ee8e4ef48b9fdb997"},{url:"/icons/maxmin.png",revision:"8f3f4c4c1f9f286c230a9b11d92a5b22"},{url:"/icons/minimize.png",revision:"70d770ab00467619d6c9a18acedae3a4"},{url:"/icons/pay_wechat.svg",revision:"2307357ff5e065b5e76b29a88aa57825"},{url:"/icons/shell_coin.svg",revision:"a910ac33d029b10d29a70c544913af7d"},{url:"/icons/stripe.svg",revision:"cba6f374e24abef406465f8bc9f69810"},{url:"/icons/token.svg",revision:"edd56c5f6887501fd36bc15176e65427"},{url:"/icons/warning.svg",revision:"3e0578143b73e97deee3bcc5140da84b"},{url:"/images/Vector.svg",revision:"99f42f921afd7e70f423b381f0f888b4"},{url:"/images/adminer.svg",revision:"dd56289e0d209796b8bace3294e37e80"},{url:"/images/allVector.svg",revision:"e8088968f740764c6d33de40d49d4605"},{url:"/images/ant-design_safety-outlined.svg",revision:"9d8b010ae9aa7ba7a5967963b1f06734"},{url:"/images/app_launchpad.svg",revision:"c8c9018141ec1b9e0453a435dd544dd4"},{url:"/images/bg-blue.svg",revision:"fb839af37ce33e4aac557aa9eb81525d"},{url:"/images/bg.svg",revision:"fef003a76670dae35667f92a631d6bab"},{url:"/images/cost_center.svg",revision:"f112a9dacd92f55051dbb29e2df711c0"},{url:"/images/database.svg",revision:"c5487ae12fb5390a2c03bd1dc6c431bc"},{url:"/images/default-user.svg",revision:"8ff96582d97a901b464bc7a907fcd1b5"},{url:"/images/docs.svg",revision:"1b291b9032ac77c49bef9b7ff2238dfe"},{url:"/images/driver-bg.png",revision:"9fab7463df083159869fc114fdb205ed"},{url:"/images/kubernetes.svg",revision:"14c4b42fd352af6d693dbdf08b8588fa"},{url:"/images/language.svg",revision:"1c98fa95b68408a038c55a6c17e5be37"},{url:"/images/lock.svg",revision:"ca1cde5a3cc61da7ee9bfb8cc7e824fe"},{url:"/images/material-symbols_expand-more-rounded.svg",revision:"07e230bf75e4aac8702222e5f2e00959"},{url:"/images/material-symbols_update.svg",revision:"e0eb0bbfae4e2e909bd87bc26a2e5d95"},{url:"/images/person.svg",revision:"dd344920d6e274ebba8e30ccb9910bdf"},{url:"/images/pgadmin.svg",revision:"b424191ad9ce93ef45a516420e506922"},{url:"/images/sealos-title.png",revision:"835ceeeb3bb84c143e18a2ff1467074c"},{url:"/images/sealos.svg",revision:"e3046b29eac7e1ef8c61d0947e7419bf"},{url:"/images/terminal.svg",revision:"7a307c5888f8d0594e121b4a361d9433"},{url:"/images/uil_info-circle.svg",revision:"7599ca28682a74ffb22fbceb6977d306"},{url:"/locales/en/cloudProviders.json",revision:"da72ba4cfed7f2e91eadf88858c3b6d6"},{url:"/locales/en/common.json",revision:"11b157752470f10e784cba00649b882b"},{url:"/locales/en/error.json",revision:"1e2d77b7465a174e01e3ec339dd6a30f"},{url:"/locales/zh/cloudProviders.json",revision:"a3fc843126ed738ccf7ce76cea5ab3b4"},{url:"/locales/zh/common.json",revision:"60bbcf50f35f9ad8aef2c4ead668282b"},{url:"/locales/zh/error.json",revision:"3b2b0b4ca63b194a5579d0baa9f28771"},{url:"/logo.svg",revision:"37f15b3541477cc046604d3d15701b4d"},{url:"/manifest.json",revision:"67647c55bd5c162f9b30078f5b200985"}],{ignoreURLParametersMatching:[]}),e.cleanupOutdatedCaches(),e.registerRoute("/",new e.NetworkFirst({cacheName:"start-url",plugins:[{cacheWillUpdate:async({request:e,response:s,event:c,state:i})=>s&&"opaqueredirect"===s.type?new Response(s.body,{status:200,statusText:"OK",headers:s.headers}):s}]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:gstatic)\.com\/.*/i,new e.CacheFirst({cacheName:"google-fonts-webfonts",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:31536e3})]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:googleapis)\.com\/.*/i,new e.StaleWhileRevalidate({cacheName:"google-fonts-stylesheets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:eot|otf|ttc|ttf|woff|woff2|font.css)$/i,new e.StaleWhileRevalidate({cacheName:"static-font-assets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:jpg|jpeg|gif|png|svg|ico|webp)$/i,new e.StaleWhileRevalidate({cacheName:"static-image-assets",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/image\?url=.+$/i,new e.StaleWhileRevalidate({cacheName:"next-image",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp3|wav|ogg)$/i,new e.CacheFirst({cacheName:"static-audio-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp4)$/i,new e.CacheFirst({cacheName:"static-video-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:js)$/i,new e.StaleWhileRevalidate({cacheName:"static-js-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:css|less)$/i,new e.StaleWhileRevalidate({cacheName:"static-style-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/data\/.+\/.+\.json$/i,new e.StaleWhileRevalidate({cacheName:"next-data",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:json|xml|csv)$/i,new e.NetworkFirst({cacheName:"static-data-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:e})=>{if(!(self.origin===e.origin))return!1;const s=e.pathname;return!s.startsWith("/api/auth/")&&!!s.startsWith("/api/")}),new e.NetworkFirst({cacheName:"apis",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:16,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:e})=>{if(!(self.origin===e.origin))return!1;return!e.pathname.startsWith("/api/")}),new e.NetworkFirst({cacheName:"others",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:e})=>!(self.origin===e.origin)),new e.NetworkFirst({cacheName:"cross-origin",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:3600})]}),"GET")}));
