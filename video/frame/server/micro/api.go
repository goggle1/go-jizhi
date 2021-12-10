package server

import (
	"log"
	"net/http"
	"time"

	// "git.tvblack.com/video/frame/core"
	// "git.tvblack.com/video/frame/server/micro/api"
	"github.com/goggle1/go-jizhi/video/frame/core"
	"github.com/goggle1/go-jizhi/video/frame/server/micro/api"

	"github.com/gorilla/mux"

	"github.com/micro/go-api/resolver"
	rrmicro "github.com/micro/go-api/resolver/micro"
	"github.com/micro/go-api/router"
	regRouter "github.com/micro/go-api/router/registry"
	aserver "github.com/micro/go-api/server"
	httpapi "github.com/micro/go-api/server/http"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/micro/api/handler"
)

type ApiIntercept func(w http.ResponseWriter, r *http.Request) bool
type ApiAfter func(w http.ResponseWriter, r *http.Request, t float64)

type srv struct {
	*mux.Router
	callback      ApiIntercept
	afterCallback ApiAfter
	upload        api.ApiUpload
}

type microApiComp struct {
	Service   micro.Service
	apiServer aserver.Server
}

func NewApiMicro(name, version, ip string, etcd []string, callback ApiIntercept, afterCallback ApiAfter, upload api.ApiUpload) *microApiComp {
	m := &microApiComp{}

	registry := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etcd
	})

	// New Service
	m.Service = micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.Registry(registry),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	h := m.getRouter(name, callback, afterCallback, upload)
	// create the server
	m.apiServer = httpapi.NewServer(ip)
	m.apiServer.Init()
	m.apiServer.Handle("/", h)

	return m
}

func (m *microApiComp) Start(app core.Appcation) {
	// Initialise service
	m.Service.Init(
		// handler wrap
		micro.WrapHandler(),
		micro.AfterStart(func() error {
			app.Run()
			return nil
		}),
		micro.AfterStop(func() error {
			app.Stop()
			return nil
		}),
	)

	app.Init()

	// Start API
	if err := m.apiServer.Start(); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := m.Service.Run(); err != nil {
		log.Fatal(err)
	}

	// Stop API
	if err := m.apiServer.Stop(); err != nil {
		log.Fatal(err)
	}
}

//{"code":100,"data":[{"detailType":1,"titleLayout":"","titleShow":1,"videoInfo":{"brief":"\u6682\u65E0\u7B80\u4ECB","img":"https:\/\/img.meiduiwang.cn\/\/pic\/xwys\/sysConfig\/value\/RM5Z2KUYXSY5A4SY8QGP.jpg","year":"2018","icon":"","blockDevices":"","title":"\u6211\u4E0D\u662F\u836F\u795E","blockArea":"","uuid":"V9HB3FP2","viewers":"","promotions":-1,"specialType":"0","act":"\u5F90\u5CE5  \u738B\u4F20\u541B  \u7AE0\u5B87  \u8C2D\u5353  \u6768\u65B0\u9E23  \u5468\u4E00\u56F4  \u738B\u4F73\u4F73  \u738B\u781A\u8F89  \u8D3E\u6668\u98DE  \u9F9A\u84D3\u82FE  \u5B81\u6D69  \u674E\u4E43\u6587","siteIds":"67AZ,P9LX","tag":"","mark":"8.4","cid":"1"},"sites":[{"site":"video","siteSkip":2,"siteId":"67AZ","siteName":"\u5176\u4ED6"},{"site":"damao","siteSkip":1,"siteId":"P9LX","siteName":"\u5927\u732B\u7F51"}],"title":"\u7535\u5F71\u8BE6\u60C5"},{"detailType":2,"titleLayout":"","titleShow":2,"volumeInfo":"","allSetsType":0,"title":"\u96C6\u6570\u5217\u8868","updateRule":""},{"detailType":4,"titleLayout":"","titleShow":1,"title":"\u76F8\u5173\u63A8\u8350","positionContent":[{"img":"http:\/\/wx4.sinaimg.cn\/mw690\/80df6fe6gy1fsp574ze9yj20u016rjyz.jpg","icon":"","title":"\u6BD2\u6218\uB3C5\uC804","uuid":"JKH4JW98","desc":"\u8BC4\u5206\uFF1A8.8"},{"img":"http:\/\/wx4.sinaimg.cn\/mw690\/80df6fe6gy1fkey56fx0qj20u0171ty1.jpg","icon":"","title":"\u771F\u5B9E\uB9AC\uC5BC","uuid":"M3BKDM3C","desc":"\u8BC4\u5206\uFF1A6.5"},{"img":"http:\/\/wx4.sinaimg.cn\/mw690\/80df6fe6gy1fkey6g3tyzj20u016r7wh.jpg","icon":"","title":"\u519B\u8230\u5C9B\uAD70\uD568\uB3C4","uuid":"MUAYYKSH","desc":"\u8BC4\u5206\uFF1A8.4"},{"img":"http:\/\/wx3.sinaimg.cn\/mw690\/80df6fe6gy1fq9da9dksij20qo15owjv.jpg","icon":"","title":"\u5BC2\u9759\u4E4B\u5730AQuietPlace","uuid":"JPLKHKRT","desc":"\u8BC4\u5206\uFF1A7.5"},{"img":"http:\/\/wx4.sinaimg.cn\/mw690\/80df6fe6gy1ft7ojbkgiij20j80rs41c.jpg","icon":"","title":"\u4E3A\u7231\u53DB\u9006\u2161\u4E4B\u9006\u6218\u5230\u5E95Baaghi2:ARebelForLove","uuid":"5HGS2EGK","desc":"\u8BC4\u5206\uFF1A7.4"},{"img":"http:\/\/wx4.sinaimg.cn\/mw690\/80df6fe6gy1ft7okcpme5j20jg0rsdjx.jpg","icon":"","title":"\u585E\u897F\u4E9A\uFF1A\u590D\u4EC7\u4E4B\u5251\u0421\u043A\u0438\u0444","uuid":"22528FSG","desc":"\u8BC4\u5206\uFF1A8.3"},{"img":"http:\/\/wx4.sinaimg.cn\/mw690\/80df6fe6gy1ft7odd581sj205006kgm4.jpg","icon":"","title":"\u7267\u91CE\u8BE1\u4E8B\u4E4B\u795E\u4ED9\u773C","uuid":"NKXLZ2YH","desc":"\u8BC4\u5206\uFF1A6.1"},{"img":"http:\/\/wx4.sinaimg.cn\/mw690\/80df6fe6gy1ft7odxzjilj205006kmxs.jpg","icon":"","title":"\u591C\u5E97\u5973\u738B2018","uuid":"ZQ9TG4G6","desc":"\u8BC4\u5206\uFF1A7.1"},{"img":"http:\/\/wx4.sinaimg.cn\/mw690\/80df6fe6gy1ft7oe09pv1j205006kaa9.jpg","icon":"","title":"\u6643\u8FC7\u4E0A\u5E1D\u4E4B\u91CD\u8FD4\u8857\u5934","uuid":"DPYJUJSG","desc":"\u8BC4\u5206\uFF1A7.5"},{"img":"http:\/\/wx4.sinaimg.cn\/mw690\/80df6fe6gy1ft6tg1rv17j20u018fwq9.jpg","icon":"","title":"\u8D1D\u9C81\u7279Beirut","uuid":"HEDSCG2K","desc":"\u8BC4\u5206\uFF1A6.8"},{"img":"http:\/\/wx4.sinaimg.cn\/mw690\/80df6fe6gy1ft6t9n5276j205k08c0t1.jpg","icon":"","title":"\u72FC\u4EBA\u795E\u63A2","uuid":"G7UV6MC5","desc":"\u8BC4\u5206\uFF1A8.5"},{"img":"http:\/\/wx4.sinaimg.cn\/mw690\/80df6fe6gy1ft6t7hd4ywj205006kdg5.jpg","icon":"","title":"\u5206\u4F539\u53F7","uuid":"AVW9L8VJ","desc":"\u8BC4\u5206\uFF1A7.7"}]}],"info":{"totalResults":3}}

func (m *microApiComp) getRouter(name string, callback ApiIntercept, afterCallback ApiAfter, upload api.ApiUpload) http.Handler {
	var h http.Handler
	r := mux.NewRouter()
	s := &srv{r, callback, afterCallback, upload}
	h = s

	r.HandleFunc("/msg", api.RPC)
	r.HandleFunc("/api", api.APIHttp)
	r.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		api.UploadMessage(w, r, upload)
	})

	// default resolve
	ropts := []resolver.Option{
		resolver.WithNamespace(name),
		resolver.WithHandler("meta"),
	}

	rr := rrmicro.NewResolver(ropts...)
	rt := regRouter.NewRouter(
		router.WithNamespace(""),
		//router.WithHandler(aapi.Handler),
		router.WithResolver(rr),
		router.WithRegistry(m.Service.Options().Registry),
	)

	r.PathPrefix("/").Handler(handler.Meta(m.Service, rt))

	//r.PathPrefix("/api").Handler(ap)
	return h
}

func (s *srv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Max-Age", "600")
	startTime := time.Now()
	if s.callback(w, r) {
		s.Router.ServeHTTP(w, r)
	}
	endTime := time.Now().Sub(startTime)
	t := endTime / 1e6
	s.afterCallback(w, r, float64(t))
}
