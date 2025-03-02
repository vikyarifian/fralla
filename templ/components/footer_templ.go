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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<footer class=\"footer footer-2\"><div class=\"footer-middle\"><div class=\"container-fluid\"><div class=\"row\"><div class=\"col-sm-12 col-lg-4\"><div class=\"widget widget-about\"><img src=\"assets/images/demos/demo-7/logo-footer.png\" class=\"footer-logo\" alt=\"Footer Logo\" width=\"105\" height=\"25\"><p>Praesent dapibus, neque id cursus ucibus, tortor neque egestas augue, eu vulputate magna eros eu erat. Aliquam erat volutpat. Nam dui mi, tincidunt quis, accumsan porttitor, facilisis luctus, metus. </p><div class=\"widget-about-info\"><div class=\"row\"><div class=\"col-sm-6 col-md-4\"><span class=\"widget-about-title\">Got Question? Call us 24/7</span> <a href=\"tel:123456789\">+0123 456 789</a></div><!-- End .col-sm-6 --><div class=\"col-sm-6 col-md-8\"><span class=\"widget-about-title\">Payment Method</span><figure class=\"footer-payments\"><img src=\"assets/images/payments.png\" alt=\"Payment methods\" width=\"272\" height=\"20\"></figure><!-- End .footer-payments --></div><!-- End .col-sm-6 --></div><!-- End .row --></div><!-- End .widget-about-info --></div><!-- End .widget about-widget --></div><!-- End .col-sm-12 col-lg-4 --><div class=\"col-sm-4 col-lg-2\"><div class=\"widget\"><h4 class=\"widget-title\">Useful links</h4><!-- End .widget-title --><ul class=\"widget-list\"><li><a href=\"about.html\">About Molla</a></li><li><a href=\"#\">How to shop on Molla</a></li><li><a href=\"faq.html\">FAQ</a></li><li><a href=\"contact.html\">Contact us</a></li><li><a href=\"login.html\">Log in</a></li></ul><!-- End .widget-list --></div><!-- End .widget --></div><!-- End .col-sm-4 col-lg-2 --><div class=\"col-sm-4 col-lg-2\"><div class=\"widget\"><h4 class=\"widget-title\">Customer Service</h4><!-- End .widget-title --><ul class=\"widget-list\"><li><a href=\"#\">Payment Methods</a></li><li><a href=\"#\">Money-back guarantee!</a></li><li><a href=\"#\">Returns</a></li><li><a href=\"#\">Shipping</a></li><li><a href=\"#\">Terms and conditions</a></li><li><a href=\"#\">Privacy Policy</a></li></ul><!-- End .widget-list --></div><!-- End .widget --></div><!-- End .col-sm-4 col-lg-2 --><div class=\"col-sm-4 col-lg-2\"><div class=\"widget\"><h4 class=\"widget-title\">My Account</h4><!-- End .widget-title --><ul class=\"widget-list\"><li><a href=\"#\">Sign In</a></li><li><a href=\"cart.html\">View Cart</a></li><li><a href=\"#\">My Wishlist</a></li><li><a href=\"#\">Track My Order</a></li><li><a href=\"#\">Help</a></li></ul><!-- End .widget-list --></div><!-- End .widget --></div><!-- End .col-sm-4 col-lg-2 --><div class=\"col-sm-6 col-lg-2\"><div class=\"widget widget-newsletter\"><h4 class=\"widget-title\">Sign up to newsletter</h4><!-- End .widget-title --><p>Aliquam erat volutpat. Nam dui mi, tincidunt quis, accumsan.</p><form action=\"#\"><div class=\"input-group\"><input type=\"email\" class=\"form-control\" placeholder=\"Enter your Email Address\" aria-label=\"Email Adress\" required><div class=\"input-group-append\"><button class=\"btn btn-dark\" type=\"submit\"><i class=\"icon-long-arrow-right\"></i></button></div><!-- .End .input-group-append --></div><!-- .End .input-group --></form></div><!-- End .widget --></div><!-- End .col-sm-6 col-lg-2 --></div><!-- End .row --></div><!-- End .container-fluid --></div><!-- End .footer-middle --><div class=\"footer-bottom\"><div class=\"container-fluid\"><p class=\"footer-copyright\">Copyright © 2019 Molla Store. All Rights Reserved.</p><!-- End .footer-copyright --><ul class=\"footer-menu\"><li><a href=\"#\">Terms Of Use</a></li><li><a href=\"#\">Privacy Policy</a></li></ul><!-- End .footer-menu --><div class=\"social-icons social-icons-color\"><span class=\"social-label\">Social Media</span> <a href=\"#\" class=\"social-icon social-facebook\" title=\"Facebook\" target=\"_blank\"><i class=\"icon-facebook-f\"></i></a> <a href=\"#\" class=\"social-icon social-twitter\" title=\"Twitter\" target=\"_blank\"><i class=\"icon-twitter\"></i></a> <a href=\"#\" class=\"social-icon social-instagram\" title=\"Instagram\" target=\"_blank\"><i class=\"icon-instagram\"></i></a> <a href=\"#\" class=\"social-icon social-youtube\" title=\"Youtube\" target=\"_blank\"><i class=\"icon-youtube\"></i></a> <a href=\"#\" class=\"social-icon social-pinterest\" title=\"Pinterest\" target=\"_blank\"><i class=\"icon-pinterest\"></i></a></div><!-- End .soial-icons --></div><!-- End .container-fluid --></div><!-- End .footer-bottom --></footer><!-- End .footer -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
