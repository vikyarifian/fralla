// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Footer() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<footer class=\"footer footer-2\" style=\"color: #777;background-color: #fff;font-weight: 200;font-size: 1.2rem;\"><div class=\"footer-middle\" style=\"padding: 2rem;\"><div class=\"container-fluid\"><div class=\"row\"><div class=\"col-sm-12 col-lg-4\"><div class=\"widget widget-about\"><p>The Right Place for All Your Shopping Needs. </p><p>Provide high-quality with affordable price. </p><p>Based in Jakarta, Indonesia.</p><div class=\"widget-about-info\"><div class=\"row\"><div class=\"col-sm-6 col-md-4\"></div><!-- End .col-sm-6 --><div class=\"col-sm-6 col-md-8\"></div><!-- End .col-sm-6 --></div><!-- End .row --></div><!-- End .widget-about-info --></div><!-- End .widget about-widget --></div><!-- End .col-sm-12 col-lg-4 --><div class=\"col-sm-4 col-lg-2\"><div class=\"widget\"><h4 class=\"widget-title\" style=\"color: #777;text-transform:none\">Useful links</h4><!-- End .widget-title --><ul class=\"widget-list\"><li><a href=\"/how-to-shop\">How to shop on Fralla</a></li><li><a href=\"/faq\">FAQ</a></li></ul><!-- End .widget-list --></div><!-- End .widget --></div><!-- End .col-sm-4 col-lg-2 --><div class=\"col-sm-4 col-lg-2\"><div class=\"widget\"><h4 class=\"widget-title\" style=\"color: #777;text-transform:none\">Customer Service</h4><!-- End .widget-title --><ul class=\"widget-list\"><li><a href=\"#\">Shipping</a></li><li><a href=\"#\">Terms and conditions</a></li><li><a href=\"#\">Privacy Policy</a></li></ul><!-- End .widget-list --></div><!-- End .widget --></div><!-- End .col-sm-4 col-lg-2 --></div><!-- End .row --></div><!-- End .container-fluid --></div><!-- End .footer-middle --><div class=\"footer-bottom\"><div class=\"container-fluid\"><p class=\"footer-copyright\">Copyright © 2025 Fralla. All Rights Reserved.</p><!-- End .footer-copyright --><div class=\"social-icons social-icons-color\"><span class=\"social-label\">Social Media</span> <a href=\"https://www.instagram.com/frallaofficial\" class=\"social-icon social-instagram\" title=\"Instagram\" target=\"_blank\"><i class=\"icon-instagram\"></i>&nbsp; Fralla Official</a> <a href=\"https://www.instagram.com/frallakids\" class=\"social-icon social-instagram\" title=\"Instagram\" target=\"_blank\"><i class=\"icon-instagram\"></i>&nbsp; Fralla Kids</a></div><!-- End .soial-icons --></div><!-- End .container-fluid --></div><!-- End .footer-bottom --></footer><!-- End .footer -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
