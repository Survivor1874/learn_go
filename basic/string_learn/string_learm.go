package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	reportId := int64(54235435432)
	fmt.Println(buildInfringementReporterNoticeSchemaUrl(reportId))
}

func buildInfringementReporterNoticeSchemaUrl(reportId int64) string {
	isObligee := "2"
	noticeType := "holder"
	if isObligee == "2" {
		noticeType = "agent"
	}
	reportIdStr := strconv.FormatInt(reportId, 10)
	var builder strings.Builder
	builder.WriteString("aweme://webview/?url=https%3A%2F%2Faweme.snssdk.com%2Ffalcon%2Fdouyin%2Finfringement_report%2F%3Fhide_nav_bar%3D1%26type%3D%")
	builder.WriteString(noticeType)
	builder.WriteString("%26report_id%3D")
	builder.WriteString(reportIdStr)
	builder.WriteString("&hide_nav_bar=1&type=")
	builder.WriteString(noticeType)
	builder.WriteString("&report_id=")
	builder.WriteString(reportIdStr)
	builder.WriteString("&rn_schema=aweme%3A%2F%2Freactnative%2F%3Fchannel%3Dfe_lynx_main_infringement_report%26bundle%3Dindex.js%26module_name%3Dpage_infringement_report%26hide_nav_bar%3D1%26loading_bgcolor%3D%2523161823%26type%3D")
	builder.WriteString(noticeType)
	builder.WriteString("%26report_id%3D")
	builder.WriteString(reportIdStr)
	return builder.String()
}
