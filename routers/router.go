package routers

import (
	"wkg/controllers/agent"
	"wkg/controllers/auth"
	"wkg/controllers/blog"
	"wkg/controllers/dash"
	"wkg/controllers/knowledge"
	"wkg/controllers/poc"
	"wkg/controllers/report"

	// "wkg/controllers/src"
	"wkg/controllers/src/company"
	"wkg/controllers/src/domain"
	"wkg/controllers/src/ip"
	"wkg/controllers/src/minipro"
	"wkg/controllers/src/notice"
	"wkg/controllers/src/service"
	"wkg/controllers/src/website"
	"wkg/controllers/src/woa"
	"wkg/controllers/task"
	"wkg/controllers/tool"
	"wkg/controllers/vuln"
	"wkg/middleware"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func InitRouter(r *gin.Engine) {

	r.Use(middleware.BruteBlock())
	r.POST("/login", auth.Login)

	v1 := r.Group("/v1")
	v1.Use(middleware.JWTAuth())
	{
		v1.GET("/getDashBoard", dash.GetDashboardInfo)
		v1.POST("/login", auth.Login)
		v1.POST("/modifyPasswd", auth.ModyfiPasswd)

		taskRouter := v1.Group("task")
		{
			taskRouter.POST("/new", task.NewTask)
			taskRouter.POST("/update", task.UpdateTask)

			taskRouter.GET("/TaskInfo", task.TaskInfo)

			taskRouter.POST("/getTaskList", task.GetTaskList)
			taskRouter.GET("/getCompIds", task.GetCompIds)
			taskRouter.GET("/getPocs", task.GetPocs)
			taskRouter.GET("/clsTaskQueue", task.ClsTaskQueue)
			taskRouter.GET("/stopTask", task.StopTask)
			taskRouter.GET("/runTask", task.RunTask)
			taskRouter.GET("/delTaskByTaskId", task.DelTaskByTaskId)
		}

		//v1.POST("/scanCompanyInfo", company.ScanCompanyInfo)
		companyRouter := v1.Group("company")
		{
			companyRouter.POST("/getCompanyInfo", company.GetCompanyInfo)
			companyRouter.POST("/newCompany", company.NewCompanyInfo)
			companyRouter.GET("/getCompanyByCId", company.GetCompanyByCId)
			// companyRouter.POST("/getCompanyInfoByKey", company.GetCompanyByKey)
			companyRouter.POST("/updateCompanyInfo", company.UpdateCompanyInfo)
			companyRouter.GET("/delCompanyByCId", company.DelCompanyByCId)

			companyRouter.GET("/getSelectOption", company.GetSelectOption)
			companyRouter.POST("/export", company.Export)
			// companyRouter.POST("/import", src.Import)
			// companyRouter.POST("/scan", src.Scan)
		}

		domainRouter := v1.Group("domain")
		{
			domainRouter.POST("/getDomainInfoByCid", domain.GetDomainInfoByCid)
			domainRouter.POST("/getNewDomainInfo", domain.GetNewDomainInfo)
			domainRouter.POST("/getDomainInfo", domain.GetDomainInfo)
			domainRouter.GET("/delDomainById", domain.DelDomainById)
			domainRouter.POST("/getDomainInfoByKey", domain.GetDomainInfoByKey)
			domainRouter.POST("/readFlagDomainInfoById", domain.ReadFlagDomainInfoById)
			domainRouter.GET("/readAllFlagDomainInfo", domain.ReadAllFlagDomainInfo)
		}

		websiteRouter := v1.Group("website")
		{
			websiteRouter.GET("/scanNew", website.ScanNewWebSite)
			websiteRouter.POST("/getWebSiteInfoByCid", website.GetWebSiteInfoByCid)
			websiteRouter.POST("/getWebSiteInfo", website.GetWebSiteInfo)
			websiteRouter.POST("/getWebSiteInfoByKey", website.GetWebSiteInfoByKey)
			websiteRouter.GET("/delWebSiteById", website.DelWebSiteById)
			websiteRouter.POST("/getNewWebSiteInfo", website.GetNewWebSiteInfo)
			websiteRouter.POST("/readFlagWebSitefoById", website.ReadFlagWebSitefoById)
			websiteRouter.GET("/readAllFlagWebSiteInfo", website.ReadAllFlagWebSiteInfo)
		}

		woaRouter := v1.Group("woa")
		{
			woaRouter.POST("/getWOAInfoByCid", woa.GetWOAInfoByCid)
			woaRouter.POST("/getWOAInfo", woa.GetWOAInfo)
			woaRouter.GET("/delWOAById", woa.DelWOAById)
			woaRouter.POST("/getNewWOAInfo", woa.GetNewWOAInfo)
			woaRouter.POST("/searchWOAInfo", woa.SearchWOAInfo)
			woaRouter.POST("/readFlagWOAInfoById", woa.ReadFlagWOAInfoById)
			woaRouter.GET("/readAllFlagWOAInfo", woa.ReadAllFlagWOAInfo)
		}

		userRouter := v1.Group("user")
		{
			userRouter.POST("/saveUser", auth.SaveUser)
			userRouter.POST("/list", auth.GetUserList)
			userRouter.GET("/del", auth.DelUser)
			userRouter.GET("/resetpwd", auth.ResetPwd)
		}

		noticeRouter := v1.Group("notice")
		{
			noticeRouter.POST("/saveNotice", notice.SaveNotice)
			noticeRouter.POST("/saveEditNotice", notice.SaveEditNotice)
			noticeRouter.POST("/getNoticeDetailById", notice.GetNoticeDetailById)
			noticeRouter.POST("/getNoticeList", notice.GetNoticeList)
			noticeRouter.GET("/delNoticeById", notice.DelNoticeById)
			noticeRouter.POST("/getNoticeByKeyword", notice.GetNoticeByKeyword)
		}

		pocRouter := v1.Group("poc")
		{
			pocRouter.POST("/getPocInfo", poc.GetPocInfo)
			pocRouter.POST("/savePoc", poc.SavePoc)
			pocRouter.POST("/saveEditPoc", poc.SaveEditPoc)
			pocRouter.POST("/getPocDetailById", poc.GetPocDetailById)
			pocRouter.POST("/getPocInfoByCid", poc.GetPocInfoByCid)
			pocRouter.GET("/delPocById", poc.DelPocById)
		}

		//v1.POST("/getNoticeInfo", notice.GetIPsInfo)
		//v1.POST("/getNewNoticeInfo", notice.GetNewIPsInfo)
		//v1.POST("/searchNoticeInfo", notice.SearchIPsInfo)
		//v1.POST("/readFlagNoticeInfoById", notice.ReadFlagIPInfoById)
		//v1.GET("/readAllFlagNoticeInfo", notice.ReadAllFlagIPInfo)

		ipRouter := v1.Group("IPs")
		{
			ipRouter.POST("/getIPsInfoByCId", ip.GetIPsInfoByCid)
			ipRouter.POST("/getIPsInfo", ip.GetIPsInfo)
			ipRouter.GET("/delIPById", ip.DelIPById)
			ipRouter.POST("/getNewIPsInfo", ip.GetNewIPsInfo)
			ipRouter.POST("/searchIPsInfo", ip.SearchIPsInfo)
			ipRouter.POST("/readFlagIPInfoById", ip.ReadFlagIPInfoById)
			ipRouter.GET("/readAllFlagIPInfo", ip.ReadAllFlagIPInfo)
		}

		miniRouter := v1.Group("mini")
		{
			miniRouter.POST("/getMiniInfoByCid", minipro.GetMiniInfoByCid)
			miniRouter.POST("/getMiniInfo", minipro.GetMiniInfo)
			miniRouter.GET("/delMiniById", minipro.DelMiniById)
			miniRouter.POST("/getNewMiniInfo", minipro.GetNewMiniInfo)
			miniRouter.POST("/searchMiniInfo", minipro.SearchMiniInfo)
			miniRouter.POST("/readFlagMiniInfoById", minipro.ReadFlagMiniInfoById)
			miniRouter.GET("/readAllFlagMiniInfo", minipro.ReadAllFlagMiniInfo)
		}

		serviceRouter := v1.Group("service")
		{
			serviceRouter.POST("/getServiceInfo", service.GetServiceInfo)
			serviceRouter.POST("/getServiceInfoByCid", service.GetServiceInfoByCid)
			serviceRouter.GET("/delServiceById", service.DelServiceById)
			serviceRouter.POST("/getNewServiceInfo", service.GetNewServiceInfo)
			serviceRouter.POST("/searchServiceInfo", service.SearchServiceInfo)
			serviceRouter.POST("/readFlagServiceInfoById", service.ReadFlagServiceInfoById)
			serviceRouter.GET("/readAllFlagServiceInfo", service.ReadAllFlagServiceInfo)
			//v1.POST("/searchServiceInfo",controllers.SearchWebSiteInfo)
		}

		bugRouter := v1.Group("bug")
		{
			bugRouter.POST("/searchOldBugReport", vuln.SearchOldBugReport)
			bugRouter.POST("/searchNewBugReport", vuln.SearchNewBugReport)
			bugRouter.POST("/searchBugCollectByKey", vuln.SearchBugCollectByKey)
			bugRouter.GET("/delBugCollectById", vuln.DelBugCollectById)
			bugRouter.POST("/getBugCollect", vuln.GetBugCollect)
			bugRouter.GET("/collect", vuln.Collect)
			bugRouter.POST("/batchRead", vuln.BatchRead)
			bugRouter.POST("/getNewBug", vuln.GetNewBug)
			bugRouter.POST("/getOldBug", vuln.GetOldBug)
			bugRouter.GET("/read", vuln.Read)
		}

		knowledgeRouter := v1.Group("knowledge")
		{

			knowledgeRouter.POST("/searchsharedKnowledgebyWord", knowledge.SearchShareKnowledgeByWord)
			knowledgeRouter.POST("/getShareknowledgeInBackend", knowledge.GetShareknowledgeInBackend)
			knowledgeRouter.GET("/cancelShare", knowledge.CancelShare)
			knowledgeRouter.GET("/cancelShareByKey", knowledge.CancelShareByKey)
			knowledgeRouter.POST("/batchCancelShare", knowledge.BatchCancelShare)

			knowledgeRouter.POST("/uploadImage", knowledge.UploadImage)

			knowledgeRouter.GET("/getSummary", knowledge.GetSummary)
			knowledgeRouter.GET("/getKnowledgeByKey", knowledge.GetKnowledgeByKey)
			knowledgeRouter.POST("/saveNewKnowledge", knowledge.SaveNewKnowledge)
			knowledgeRouter.POST("/saveEditKnowledge", knowledge.SaveEditKnowledge)
			knowledgeRouter.GET("/getTree", knowledge.GetTree)
			knowledgeRouter.GET("/addNode", knowledge.AddNode)
			knowledgeRouter.POST("/delTreeNode", knowledge.DelTreeNode)
			knowledgeRouter.POST("/modifyTreeNode", knowledge.ModifyTreeNode)
			knowledgeRouter.GET("/shareknowledge", knowledge.Shareknowledge)
			knowledgeRouter.GET("/getKnowledgeCategories", knowledge.GetKnowledgeCategories)
		}

		// blogRouter := v1.Group("blog")
		// {

		//v1.GET("/getKnowledge",controllers.GetKnowledge)

		// v1.GET("/delTag", knowledge.DelTag)
		// v1.GET("/addTag", knowledge.AddTag)
		// }

		agentRouter := v1.Group("agent")
		{
			agentRouter.POST("/list", agent.AgentList)
			agentRouter.GET("/delete", agent.DeleteAgent)
			agentRouter.POST("/getAgentById", agent.GetAgentById)
			agentRouter.GET("/AgentDetail", agent.AgentDetail)
		}

	}

	//beego.Router("/debug/pprof", &controllers.ProfController{})
	//beego.Router("/debug/pprof/:app([\\w]+)", &controllers.ProfController{})

	v2 := r.Group("/v2")
	{

		v2.GET("/getHotList", knowledge.GetHotList)
		v2.POST("/report/upload", report.Upload)
		v2.POST("/agent/heartBeat", agent.HeartBeat)
		//v2.GET("/domainScan", controllers.DomainScan)
		//v2.GET("/VulnScanSingle",controllers.VulnScanSingle)
		//v2.GET("/VulnScanMulti",controllers.VulnScanMulti)

		v2.GET("/getRootDomain", vuln.GetRootDomain)
		v2.GET("/getFofaRule", vuln.GetFofaRule)
		v2.POST("/bugReport", vuln.BugReport)

	}

	blogRouter := v2.Group("blog")
	{
		blogRouter.POST("/getsharedKnowledgeListbyTag", blog.GetsharedKnowledgeListbyTag)
		blogRouter.GET("/getShareknowledgeByKey", blog.GetShareknowledgeByKey)
		blogRouter.POST("/getShareknowledgeList", blog.GetShareknowledgeList)
		blogRouter.POST("/searchsharedKnowledgebyWord", blog.SearchShareKnowledgeByWord)
		blogRouter.GET("/getTags", blog.GetTags)
	}

	toolRouter := v2.Group("tool")
	{
		toolRouter.POST("/sendQuestion", tool.SendQuestion)
	}

	v3 := r.Group("/v3")
	v3.Use(middleware.InputToken())
	{
		// v3.POST("/import", src.Import)
		v3.GET("/task/getTask", task.GetTask)
		v3.POST("/task/uploadDomains", task.UploadDomains)

	}
}
