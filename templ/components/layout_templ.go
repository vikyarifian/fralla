// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fralla/dto"

func Layout(title string, token dto.Token) templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><script src=\"https://unpkg.com/htmx.org@2.0.1\" integrity=\"sha384-QWGpdj554B4ETpJJC9z+ZHJcA/i59TyjxEPXiiUgN2WmTyV5OEZWCD6gQhgkdpB/\" crossorigin=\"anonymous\"></script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Head(title).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<style>\n            @media screen and (max-width: 480px) {\n                .banner_fp1 {\n                    visibility: hidden;\n                    display: none;\n                }\n            }\n            @media screen and (max-width: 480px) {\n                .banner_fp2 {\n                    visibility: show;\n                    display: block;\n                }\n            }\n            @media screen and (min-width: 480px) {\n                .banner_fp2 {\n                    visibility: hidden;\n                    display: none;\n                }\n                .header-7 {\n                    padding-left: 4rem;\n                }\n            }\n            @media screen and (max-width: 1000px) {\n                .header-7 {\n                    padding-left: 1rem;\n                }\n            }\n            @media screen and (max-width: 1000px) {\n                .header-7 {\n                    padding-left: 1rem;\n                }\n            }\n            /* Absolute Center Spinner */\n\t\t\t.screenloading {\n\t\t\t\tposition: fixed;\n\t\t\t\tz-index: 999999;\n\t\t\t\theight: 2em;\n\t\t\t\twidth: 2em;\n\t\t\t\toverflow: show;\n\t\t\t\tmargin: auto;\n\t\t\t\ttop: 0;\n\t\t\t\tleft: 0;\n\t\t\t\tbottom: 0;\n\t\t\t\tright: 0;\n\t\t\t}\n\n\t\t\t/* Transparent Overlay */\n\t\t\t.screenloading:before {\n\t\t\t\tcontent: '';\n\t\t\t\tdisplay: block;\n\t\t\t\tposition: fixed;\n\t\t\t\ttop: 0;\n\t\t\t\tleft: 0;\n\t\t\t\twidth: 100%;\n\t\t\t\theight: 100%;\n\t\t\t\tbackground: radial-gradient(rgba(122, 120, 120, 0.8), rgba(122, 120, 120, 0.8));\n\t\t\t\tbackground: -webkit-radial-gradient(rgba(122, 120, 120, 0.8), rgba(122, 120, 120, 0.8));\n\t\t\t\topacity: 0.8;\n\t\t\t}\n\n\t\t\t/* :not(:required) hides these rules from IE9 and below */\n\t\t\t.screenloading:not(:required) {\n\t\t\t\t/* hide \"loading...\" text */\n\t\t\t\tfont: 0/0 a;\n\t\t\t\tcolor: transparent;\n\t\t\t\ttext-shadow: none;\n\t\t\t\tbackground-color: transparent;\n\t\t\t\tborder: 0;\n\t\t\t}\n\n\t\t\t.screenloading:not(:required):after {\n\t\t\t\tcontent: '';\n\t\t\t\tdisplay: block;\n\t\t\t\tfont-size: 10px;\n\t\t\t\twidth: 1em;\n\t\t\t\theight: 1em;\n\t\t\t\tmargin-top: -0.5em;\n\t\t\t\t-webkit-animation: spinner 150ms infinite linear;\n\t\t\t\t-moz-animation: spinner 150ms infinite linear;\n\t\t\t\t-ms-animation: spinner 150ms infinite linear;\n\t\t\t\t-o-animation: spinner 150ms infinite linear;\n\t\t\t\tanimation: spinner 150ms infinite linear;\n\t\t\t\tborder-radius: 0.5em;\n\t\t\t\t-webkit-box-shadow: rgba(255,255,255, 0.75) 1.5em 0 0 0, rgba(255,255,255, 0.75) 1.1em 1.1em 0 0, rgba(255,255,255, 0.75) 0 1.5em 0 0, rgba(255,255,255, 0.75) -1.1em 1.1em 0 0, rgba(255,255,255, 0.75) -1.5em 0 0 0, rgba(255,255,255, 0.75) -1.1em -1.1em 0 0, rgba(255,255,255, 0.75) 0 -1.5em 0 0, rgba(255,255,255, 0.75) 1.1em -1.1em 0 0;\n\t\t\t\tbox-shadow: rgba(255,255,255, 0.75) 1.5em 0 0 0, rgba(255,255,255, 0.75) 1.1em 1.1em 0 0, rgba(255,255,255, 0.75) 0 1.5em 0 0, rgba(255,255,255, 0.75) -1.1em 1.1em 0 0, rgba(255,255,255, 0.75) -1.5em 0 0 0, rgba(255,255,255, 0.75) -1.1em -1.1em 0 0, rgba(255,255,255, 0.75) 0 -1.5em 0 0, rgba(255,255,255, 0.75) 1.1em -1.1em 0 0;\n\t\t\t\topacity: 1;\n\t\t\t}\n\n\t\t\t/* Animation */\n\n\t\t\t@-webkit-keyframes spinner {\n\t\t\t\t0% {\n\t\t\t\t\t-webkit-transform: rotate(0deg);\n\t\t\t\t\t-moz-transform: rotate(0deg);\n\t\t\t\t\t-ms-transform: rotate(0deg);\n\t\t\t\t\t-o-transform: rotate(0deg);\n\t\t\t\t\ttransform: rotate(0deg);\n\t\t\t\t}\n\t\t\t\t100% {\n\t\t\t\t\t-webkit-transform: rotate(360deg);\n\t\t\t\t\t-moz-transform: rotate(360deg);\n\t\t\t\t\t-ms-transform: rotate(360deg);\n\t\t\t\t\t-o-transform: rotate(360deg);\n\t\t\t\t\ttransform: rotate(360deg);\n\t\t\t\t}\n\t\t\t}\n\t\t\t@-moz-keyframes spinner {\n\t\t\t\t0% {\n\t\t\t\t\t-webkit-transform: rotate(0deg);\n\t\t\t\t\t-moz-transform: rotate(0deg);\n\t\t\t\t\t-ms-transform: rotate(0deg);\n\t\t\t\t\t-o-transform: rotate(0deg);\n\t\t\t\t\ttransform: rotate(0deg);\n\t\t\t\t}\n\t\t\t\t100% {\n\t\t\t\t\t-webkit-transform: rotate(360deg);\n\t\t\t\t\t-moz-transform: rotate(360deg);\n\t\t\t\t\t-ms-transform: rotate(360deg);\n\t\t\t\t\t-o-transform: rotate(360deg);\n\t\t\t\t\ttransform: rotate(360deg);\n\t\t\t\t}\n\t\t\t}\n\t\t\t@-o-keyframes spinner {\n\t\t\t\t0% {\n\t\t\t\t\t-webkit-transform: rotate(0deg);\n\t\t\t\t\t-moz-transform: rotate(0deg);\n\t\t\t\t\t-ms-transform: rotate(0deg);\n\t\t\t\t\t-o-transform: rotate(0deg);\n\t\t\t\t\ttransform: rotate(0deg);\n\t\t\t\t}\n\t\t\t\t100% {\n\t\t\t\t\t-webkit-transform: rotate(360deg);\n\t\t\t\t\t-moz-transform: rotate(360deg);\n\t\t\t\t\t-ms-transform: rotate(360deg);\n\t\t\t\t\t-o-transform: rotate(360deg);\n\t\t\t\t\ttransform: rotate(360deg);\n\t\t\t\t}\n\t\t\t}\n\t\t\t@keyframes spinner {\n\t\t\t\t0% {\n\t\t\t\t\t-webkit-transform: rotate(0deg);\n\t\t\t\t\t-moz-transform: rotate(0deg);\n\t\t\t\t\t-ms-transform: rotate(0deg);\n\t\t\t\t\t-o-transform: rotate(0deg);\n\t\t\t\t\ttransform: rotate(0deg);\n\t\t\t\t}\n\t\t\t\t100% {\n\t\t\t\t\t-webkit-transform: rotate(360deg);\n\t\t\t\t\t-moz-transform: rotate(360deg);\n\t\t\t\t\t-ms-transform: rotate(360deg);\n\t\t\t\t\t-o-transform: rotate(360deg);\n\t\t\t\t\ttransform: rotate(360deg);\n\t\t\t\t}\n\t\t\t}\n\n            .htmx-loader {\n\t\t\t\tdisplay: none;\n\t\t\t}\n\n\t\t\t.htmx-request .htmx-loader {\n\t\t\t\tdisplay: block;\n\t\t\t}\n\n\t\t\t.htmx-request.htmx-loader {\n\t\t\t\tdisplay: block;\n\t\t\t}\n\n\t\t\t/* Hide previously loaded content during HTMX request */\n\t\t\t.htmx-request .loaded-content {\n\t\t\t\tdisplay: none;\n\t\t\t}\n\n\t\t\t.htmx-request.loaded-content {\n\t\t\t\tdisplay: none;\n\t\t\t}\n\n\t\t\t.scoll-pane {\n\t\t\t\toverflow: auto;\n\t\t\t\twidth: 40vw;\n\t\t\t\tmax-height: 32vh;\n\t\t\t\tbackground: hsl(0, 0%, 100%);\n\t\t\t\tpadding: 1rem;\n\t\t\t\tborder-radius: 0.5rem;\n\t\t\t}\n\n\t\t\t::-webkit-scrollbar {\n\t\t\t\twidth: 12px; /* Mostly for vertical scrollbars */\n\t\t\t\theight: 12px; /* Mostly for horizontal scrollbars */\n\t\t\t}\n\t\t\t::-webkit-scrollbar-thumb { /* Foreground */\n\t\t\t\tbackground: var(--scrollbar-foreground);\n\t\t\t\tbackground: rgb(255, 255, 255);\n\t\t\t\tborder-radius: 999px;\n\t\t\t\tborder: 3px solid transparent;\n\t\t\t\tbackground-clip: padding-box;\n\t\t\t}\n\t\t\t::-webkit-scrollbar-track { /* Background */\n\t\t\t\tbackground: var(--scrollbar-background);\n\t\t\t\tbackground: transparent;\n\t\t\t}\n        </style><body style=\"font: normal 200 1.2rem / 1.46 &#39;Poppins&#39;, sans-serif;\"><div class=\"page-wrapper\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Header(token).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<main class=\"main\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<div class=\"screenloading htmx-loader htmx-indicator\">Loading&#8230;</div></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Footer().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</div><button id=\"scroll-top\" title=\"Back to Top\"><i class=\"icon-arrow-up\"></i></button><!-- Mobile Menu --><div class=\"mobile-menu-overlay\"></div><!-- End .mobil-menu-overlay --><div class=\"mobile-menu-container\"><div class=\"mobile-menu-wrapper\"><span class=\"mobile-menu-close\"><i class=\"icon-close\"></i></span><form action=\"#\" method=\"get\" class=\"mobile-search\"><label for=\"mobile-search\" class=\"sr-only\">Search</label> <input type=\"search\" class=\"form-control\" name=\"mobile-search\" id=\"mobile-search\" placeholder=\"Search product...\" required> <button class=\"btn btn-primary\" type=\"submit\"><i class=\"icon-search\"></i></button></form><nav class=\"mobile-nav\"><ul class=\"mobile-menu\"><li><a href=\"/shop\" style=\"text-transform:none\">Shop</a></li><li><a href=\"#\" style=\"text-transform:none\">About</a></li><li><a href=\"#\" style=\"text-transform:none\">Contact</a></li></ul></nav><!-- End .mobile-nav --></div><!-- End .mobile-menu-wrapper --></div><!-- End .mobile-menu-container -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !token.IsAuth {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<!-- Sign in / Register Modal --> <div class=\"modal fade\" id=\"signin-modal\" tabindex=\"-1\" role=\"dialog\" aria-hidden=\"true\"><div class=\"modal-dialog modal-dialog-centered\" role=\"document\"><div class=\"modal-content\"><div class=\"modal-body\"><button type=\"button\" class=\"close\" data-dismiss=\"modal\" aria-label=\"Close\"><span aria-hidden=\"true\"><i class=\"icon-close\"></i></span></button><div class=\"form-box\"><div class=\"form-tab\"><ul class=\"nav nav-pills nav-fill\" role=\"tablist\"><li class=\"nav-item\"><a class=\"nav-link active\" id=\"signin-tab\" data-toggle=\"tab\" href=\"#signin\" role=\"tab\" aria-controls=\"signin\" aria-selected=\"true\">Sign In</a></li><li class=\"nav-item\"><a class=\"nav-link\" id=\"register-tab\" data-toggle=\"tab\" href=\"#register\" role=\"tab\" aria-controls=\"register\" aria-selected=\"false\">Register</a></li></ul><div class=\"tab-content\" id=\"tab-content-5\"><div class=\"tab-pane fade show active\" id=\"signin\" role=\"tabpanel\" aria-labelledby=\"signin-tab\"><form id=\"login-form\" hx-disabled-elt=\"#login-button\" hx-post=\"/login\" hx-swap=\"none\" hx-on::before-request=\"document.getElementById(&#39;error_login&#39;).innerHTML = &#39;&#39;\" hx-target-400=\"#error_login\" hx-indicator=\".htmx-loader\"><div class=\"form-group\"><label for=\"singin-email_phone\">Phone or email address *</label> <input type=\"text\" class=\"form-control\" id=\"singin-email_phone\" name=\"email_phone\" required></div><!-- End .form-group --><div class=\"form-group\"><label for=\"singin-password\">Password *</label> <input type=\"password\" class=\"form-control\" id=\"singin-password\" name=\"password\" required></div><!-- End .form-group --><span id=\"error_login\"></span> <span id=\"success_login\"></span><div class=\"form-footer\"><button id=\"login-button\" type=\"submit\" class=\"btn btn-outline-primary-2\"><span>LOG IN</span> <i class=\"icon-long-arrow-right\"></i></button><div class=\"custom-control custom-checkbox\"><input type=\"checkbox\" class=\"custom-control-input\" id=\"signin-remember\"> <label class=\"custom-control-label\" for=\"signin-remember\">Remember Me</label></div><!-- End .custom-checkbox --><a href=\"#\" class=\"forgot-link\">Forgot Your Password?</a></div><!-- End .form-footer --></form><div class=\"form-choice\"><p class=\"text-center\">or sign in with</p><div class=\"row\"><div class=\"col-sm-6\"><a href=\"#\" class=\"btn btn-login btn-g\"><i class=\"icon-google\"></i> Login With Google</a></div><!-- End .col-6 --><div class=\"col-sm-6\"><a href=\"#\" class=\"btn btn-login btn-f\"><i class=\"icon-facebook-f\"></i> Login With Facebook</a></div><!-- End .col-6 --></div><!-- End .row --></div><!-- End .form-choice --></div><!-- .End .tab-pane --><div class=\"tab-pane fade\" id=\"register\" role=\"tabpanel\" aria-labelledby=\"register-tab\"><form id=\"register-form\" hx-disabled-elt=\"#register-button\" hx-post=\"/register\" hx-swap=\"none\" hx-on::before-request=\"document.getElementById(&#39;error_register&#39;).innerHTML = &#39;&#39;;document.getElementById(&#39;success_register&#39;).innerHTML = &#39;&#39;\" hx-on::after-request=\"if(event.detail.xhr.status==200) {\n                                                                                document.getElementById(&#39;register-email&#39;).value = &#39;&#39;;document.getElementById(&#39;register-phone&#39;).value = &#39;&#39;;\n                                                                                document.getElementById(&#39;register-password&#39;).value = &#39;&#39;;\n                                                                            }\" hx-target-400=\"#error_register\" hx-indicator=\".htmx-loader\"><div class=\"form-group\"><label for=\"register-email\">Your email address *</label> <input type=\"email\" class=\"form-control\" id=\"register-email\" name=\"email\" required></div><!-- End .form-group --><div class=\"form-group\"><label for=\"register-phone\">Your phone number*</label> <input type=\"phone\" class=\"form-control\" id=\"register-phone\" name=\"phone\" required></div><!-- End .form-group --><div class=\"form-group\"><label for=\"register-password\">Password *</label> <input type=\"password\" class=\"form-control\" id=\"register-password\" name=\"password\" required></div><!-- End .form-group --><div class=\"htmx-indicator\"></div><span id=\"error_register\"></span> <span id=\"success_register\"></span><div class=\"form-footer\"><button id=\"register-button\" type=\"submit\" class=\"btn btn-outline-primary-2\"><span>SIGN UP</span> <i class=\"icon-long-arrow-right\"></i></button><div class=\"custom-control custom-checkbox\"><input type=\"checkbox\" class=\"custom-control-input\" id=\"register-policy\" required> <label class=\"custom-control-label\" for=\"register-policy\">I agree to the <a href=\"#\">privacy policy</a> *</label></div><!-- End .custom-checkbox --></div><!-- End .form-footer --></form><div class=\"form-choice\"><p class=\"text-center\">or sign in with</p><div class=\"row\"><div class=\"col-sm-6\"><a href=\"#\" class=\"btn btn-login btn-g\"><i class=\"icon-google\"></i> Login With Google</a></div><!-- End .col-6 --><div class=\"col-sm-6\"><a href=\"#\" class=\"btn btn-login  btn-f\"><i class=\"icon-facebook-f\"></i> Login With Facebook</a></div><!-- End .col-6 --></div><!-- End .row --></div><!-- End .form-choice --></div><!-- .End .tab-pane --></div><!-- End .tab-content --></div><!-- End .form-tab --></div><!-- End .form-box --></div><!-- End .modal-body --></div><!-- End .modal-content --></div><!-- End .modal-dialog --></div><!-- End .modal -->")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if token.PopUpShow {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<div class=\"container newsletter-popup-container mfp-hide\" id=\"newsletter-popup-form\"><div class=\"row justify-content-center\"><div class=\"col-10\"><div class=\"row no-gutters bg-white newsletter-popup-content\"><div class=\"col-xl-3-5col col-lg-7 banner-content-wrap\"><div class=\"banner-content text-center\"><img src=\"assets/images/popup/newsletter/logo-2.png\" class=\"logo\" alt=\"logo\" width=\"60\" height=\"15\"><h2 class=\"banner-title\">Don't <span>miss out<light>!</light></span></h2><p>Join Fralla group whatsapp.</p><form action=\"#\"><div class=\"input-group input-group-round\"><a href=\"https://wa.me/6282128905833\" class=\"input-group-append col-sm-12\" target=\"_blank\"><div class=\"col-sm-8\"><svg height=\"40%\" style=\"fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:2;\" version=\"1.1\" viewBox=\"0 0 512 512\" width=\"40%\" xml:space=\"preserve\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:serif=\"http://www.serif.com/\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"><g id=\"WhatsApp-Logo-Icon\"><path d=\"M116.225,-0.001c-11.264,0.512 -26.112,1.536 -32.768,3.072c-10.24,2.048 -19.968,5.12 -27.648,9.216c-9.728,4.608 -17.92,10.752 -25.088,17.92c-7.68,7.68 -13.824,15.872 -18.432,25.6c-4.096,7.68 -7.168,17.408 -9.216,27.648c-1.536,6.656 -2.56,21.504 -2.56,32.768c-0.512,4.608 -0.512,10.752 -0.512,13.824l0,251.905l0,13.824c0.512,11.264 1.536,26.112 3.072,32.768c2.048,10.24 5.12,19.968 9.216,27.648c4.608,9.728 10.752,17.92 17.92,25.088c7.68,7.68 15.872,13.824 25.6,18.432c7.68,4.096 17.408,7.168 27.648,9.216c6.656,1.536 21.504,2.56 32.768,2.56c4.608,0.512 10.752,0.512 13.824,0.512l251.904,0l13.824,0c11.264,-0.512 26.112,-1.536 32.768,-3.072c10.24,-2.048 19.968,-5.12 27.648,-9.216c9.728,-4.608 17.92,-10.752 25.088,-17.92c7.68,-7.68 13.824,-15.872 18.432,-25.6c4.096,-7.68 7.168,-17.408 9.216,-27.648c1.536,-6.656 2.56,-21.504 2.56,-32.768c0.512,-4.608 0.512,-10.752 0.512,-13.824l0,-265.729c-0.512,-11.264 -1.536,-26.112 -3.072,-32.768c-2.048,-10.24 -5.12,-19.968 -9.216,-27.648c-4.608,-9.728 -10.752,-17.92 -17.92,-25.088c-7.68,-7.68 -15.872,-13.824 -25.6,-18.432c-7.68,-4.096 -17.408,-7.168 -27.648,-9.216c-6.656,-1.536 -21.504,-2.56 -32.768,-2.56c-4.608,-0.512 -10.752,-0.512 -13.824,-0.512l-265.728,0Z\" style=\"fill:url(#_Linear1);fill-rule:nonzero;\"></path><path d=\"M344.754,289.698c-4.56,-2.282 -26.98,-13.311 -31.161,-14.832c-4.18,-1.521 -7.219,-2.282 -10.259,2.282c-3.041,4.564 -11.78,14.832 -14.44,17.875c-2.66,3.042 -5.32,3.423 -9.88,1.14c-4.561,-2.281 -19.254,-7.095 -36.672,-22.627c-13.556,-12.087 -22.709,-27.017 -25.369,-31.581c-2.66,-4.564 -0.283,-7.031 2,-9.304c2.051,-2.041 4.56,-5.324 6.84,-7.986c2.28,-2.662 3.04,-4.564 4.56,-7.606c1.52,-3.042 0.76,-5.705 -0.38,-7.987c-1.14,-2.282 -10.26,-24.72 -14.06,-33.848c-3.701,-8.889 -7.461,-7.686 -10.26,-7.826c-2.657,-0.132 -5.7,-0.16 -8.74,-0.16c-3.041,0 -7.98,1.141 -12.161,5.704c-4.18,4.564 -15.96,15.594 -15.96,38.032c0,22.438 16.34,44.116 18.62,47.159c2.281,3.043 32.157,49.089 77.902,68.836c10.88,4.697 19.374,7.501 25.997,9.603c10.924,3.469 20.866,2.98 28.723,1.806c8.761,-1.309 26.98,-11.029 30.781,-21.677c3.799,-10.649 3.799,-19.777 2.659,-21.678c-1.139,-1.902 -4.179,-3.043 -8.74,-5.325m-83.207,113.573l-0.061,0c-27.22,-0.011 -53.917,-7.32 -77.207,-21.137l-5.539,-3.287l-57.413,15.056l15.325,-55.959l-3.608,-5.736c-15.184,-24.145 -23.203,-52.051 -23.192,-80.704c0.033,-83.611 68.083,-151.635 151.756,-151.635c40.517,0.016 78.603,15.811 107.243,44.474c28.64,28.663 44.404,66.764 44.389,107.283c-0.035,83.617 -68.083,151.645 -151.693,151.645m129.102,-280.709c-34.457,-34.486 -80.281,-53.487 -129.103,-53.507c-100.595,0 -182.468,81.841 -182.508,182.437c-0.013,32.156 8.39,63.546 24.361,91.212l-25.892,94.545l96.75,-25.37c26.657,14.535 56.67,22.194 87.216,22.207l0.075,0c100.586,0 182.465,-81.852 182.506,-182.448c0.019,-48.751 -18.946,-94.59 -53.405,-129.076\" style=\"fill:#fff;\"></path></g><defs><linearGradient gradientTransform=\"matrix(0,-512,-512,0,256.001,512)\" gradientUnits=\"userSpaceOnUse\" id=\"_Linear1\" x1=\"0\" x2=\"1\" y1=\"0\" y2=\"0\"><stop offset=\"0\" style=\"stop-color:#25cf43;stop-opacity:1\"></stop><stop offset=\"1\" style=\"stop-color:#61fd7d;stop-opacity:1\"></stop></linearGradient></defs></svg> Chat Fralla Here</div></a><!-- .End .input-group-append --><a href=\"https://chat.whatsapp.com/CNxz2oS4btbFTwhn2qPzE0\" class=\"input-group-append col-sm-12\" target=\"_blank\"><div class=\"col-sm-8\"><svg height=\"40%\" style=\"fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:2;\" version=\"1.1\" viewBox=\"0 0 512 512\" width=\"40%\" xml:space=\"preserve\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:serif=\"http://www.serif.com/\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"><g id=\"WhatsApp-Logo-Icon\"><path d=\"M116.225,-0.001c-11.264,0.512 -26.112,1.536 -32.768,3.072c-10.24,2.048 -19.968,5.12 -27.648,9.216c-9.728,4.608 -17.92,10.752 -25.088,17.92c-7.68,7.68 -13.824,15.872 -18.432,25.6c-4.096,7.68 -7.168,17.408 -9.216,27.648c-1.536,6.656 -2.56,21.504 -2.56,32.768c-0.512,4.608 -0.512,10.752 -0.512,13.824l0,251.905l0,13.824c0.512,11.264 1.536,26.112 3.072,32.768c2.048,10.24 5.12,19.968 9.216,27.648c4.608,9.728 10.752,17.92 17.92,25.088c7.68,7.68 15.872,13.824 25.6,18.432c7.68,4.096 17.408,7.168 27.648,9.216c6.656,1.536 21.504,2.56 32.768,2.56c4.608,0.512 10.752,0.512 13.824,0.512l251.904,0l13.824,0c11.264,-0.512 26.112,-1.536 32.768,-3.072c10.24,-2.048 19.968,-5.12 27.648,-9.216c9.728,-4.608 17.92,-10.752 25.088,-17.92c7.68,-7.68 13.824,-15.872 18.432,-25.6c4.096,-7.68 7.168,-17.408 9.216,-27.648c1.536,-6.656 2.56,-21.504 2.56,-32.768c0.512,-4.608 0.512,-10.752 0.512,-13.824l0,-265.729c-0.512,-11.264 -1.536,-26.112 -3.072,-32.768c-2.048,-10.24 -5.12,-19.968 -9.216,-27.648c-4.608,-9.728 -10.752,-17.92 -17.92,-25.088c-7.68,-7.68 -15.872,-13.824 -25.6,-18.432c-7.68,-4.096 -17.408,-7.168 -27.648,-9.216c-6.656,-1.536 -21.504,-2.56 -32.768,-2.56c-4.608,-0.512 -10.752,-0.512 -13.824,-0.512l-265.728,0Z\" style=\"fill:url(#_Linear1);fill-rule:nonzero;\"></path><path d=\"M344.754,289.698c-4.56,-2.282 -26.98,-13.311 -31.161,-14.832c-4.18,-1.521 -7.219,-2.282 -10.259,2.282c-3.041,4.564 -11.78,14.832 -14.44,17.875c-2.66,3.042 -5.32,3.423 -9.88,1.14c-4.561,-2.281 -19.254,-7.095 -36.672,-22.627c-13.556,-12.087 -22.709,-27.017 -25.369,-31.581c-2.66,-4.564 -0.283,-7.031 2,-9.304c2.051,-2.041 4.56,-5.324 6.84,-7.986c2.28,-2.662 3.04,-4.564 4.56,-7.606c1.52,-3.042 0.76,-5.705 -0.38,-7.987c-1.14,-2.282 -10.26,-24.72 -14.06,-33.848c-3.701,-8.889 -7.461,-7.686 -10.26,-7.826c-2.657,-0.132 -5.7,-0.16 -8.74,-0.16c-3.041,0 -7.98,1.141 -12.161,5.704c-4.18,4.564 -15.96,15.594 -15.96,38.032c0,22.438 16.34,44.116 18.62,47.159c2.281,3.043 32.157,49.089 77.902,68.836c10.88,4.697 19.374,7.501 25.997,9.603c10.924,3.469 20.866,2.98 28.723,1.806c8.761,-1.309 26.98,-11.029 30.781,-21.677c3.799,-10.649 3.799,-19.777 2.659,-21.678c-1.139,-1.902 -4.179,-3.043 -8.74,-5.325m-83.207,113.573l-0.061,0c-27.22,-0.011 -53.917,-7.32 -77.207,-21.137l-5.539,-3.287l-57.413,15.056l15.325,-55.959l-3.608,-5.736c-15.184,-24.145 -23.203,-52.051 -23.192,-80.704c0.033,-83.611 68.083,-151.635 151.756,-151.635c40.517,0.016 78.603,15.811 107.243,44.474c28.64,28.663 44.404,66.764 44.389,107.283c-0.035,83.617 -68.083,151.645 -151.693,151.645m129.102,-280.709c-34.457,-34.486 -80.281,-53.487 -129.103,-53.507c-100.595,0 -182.468,81.841 -182.508,182.437c-0.013,32.156 8.39,63.546 24.361,91.212l-25.892,94.545l96.75,-25.37c26.657,14.535 56.67,22.194 87.216,22.207l0.075,0c100.586,0 182.465,-81.852 182.506,-182.448c0.019,-48.751 -18.946,-94.59 -53.405,-129.076\" style=\"fill:#fff;\"></path></g><defs><linearGradient gradientTransform=\"matrix(0,-512,-512,0,256.001,512)\" gradientUnits=\"userSpaceOnUse\" id=\"_Linear1\" x1=\"0\" x2=\"1\" y1=\"0\" y2=\"0\"><stop offset=\"0\" style=\"stop-color:#25cf43;stop-opacity:1\"></stop><stop offset=\"1\" style=\"stop-color:#61fd7d;stop-opacity:1\"></stop></linearGradient></defs></svg> Join Group Here</div></a><!-- .End .input-group-append --></div><!-- .End .input-group --></form><div class=\"custom-control custom-checkbox\"><input type=\"checkbox\" class=\"custom-control-input\" id=\"register-policy-2\" name=\"popup\" hx-get=\"/hide-popup\" hx-swap=\"none\" hx-trigger=\"click\"> <label class=\"custom-control-label\" for=\"register-policy-2\">Do not show this popup again</label></div><!-- End .custom-checkbox --></div></div><div class=\"col-xl-2-5col col-lg-5 \"><img src=\"assets/images/popup/newsletter/img-3.jpg\" class=\"newsletter-img\" alt=\"newsletter\"></div></div></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<!-- Plugins JS File --><script src=\"assets/js/jquery.min.js\"></script><script src=\"assets/js/bootstrap.bundle.min.js\"></script><script src=\"assets/js/jquery.hoverIntent.min.js\"></script><script src=\"assets/js/jquery.waypoints.min.js\"></script><script src=\"assets/js/superfish.min.js\"></script><script src=\"assets/js/bootstrap-input-spinner.js\"></script><script src=\"assets/js/owl.carousel.min.js\"></script><script src=\"assets/js/jquery.plugin.min.js\"></script><script src=\"assets/js/jquery.magnific-popup.min.js\"></script><script src=\"assets/js/jquery.countdown.min.js\"></script><!-- Main JS File --><script src=\"assets/js/main.js\"></script><script src=\"assets/js/demos/demo-7.js\"></script></body><script>\n            document.addEventListener(\"DOMContentLoaded\", (event) => {\n                document.body.addEventListener('htmx:beforeSwap', function(evt) {\n                    if (evt.detail.xhr.status === 400) {\n                        evt.detail.shouldSwap = true;\n                        evt.detail.isError = false;\n                    }\n                });\n            })\t\n        </script></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
