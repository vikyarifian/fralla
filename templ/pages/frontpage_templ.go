// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fralla/dto"
	"fralla/templ/components"
)

func FrontPage(token dto.Token) templ.Component {
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"container-fluid\"><div class=\"row\"><div class=\"col-lg-6\"><div class=\"banner banner-big banner-overlay\"><a href=\"#\"><img src=\"assets/images/demos/demo-7/banners/banner-6.jpg\" alt=\"Banner\"></a><div class=\"banner-content banner-content-center\"><h3 class=\"banner-subtitle text-white\"><a href=\"#\">New Collection</a></h3><!-- End .banner-subtitle --><h2 class=\"banner-title text-white\"><a href=\"#\">Shop Kid's</a></h2><!-- End .banner-title --><a href=\"#\" class=\"btn underline\"><span>Discover Now</span></a></div><!-- End .banner-content --></div><!-- End .banner --></div><!-- End .col-lg-6 --><div class=\"col-lg-6\"><div class=\"banner banner-big banner-overlay\"><a href=\"#\"><img src=\"assets/images/demos/demo-7/banners/banner-7.jpg\" alt=\"Banner\"></a><div class=\"banner-content banner-content-center\"><h3 class=\"banner-subtitle text-white\"><a href=\"#\">New Collection</a></h3><!-- End .banner-subtitle --><h2 class=\"banner-title text-white\"><a href=\"#\">Shop Women's</a></h2><!-- End .banner-title --><a href=\"#\" class=\"btn underline\"><span>Discover Now</span></a></div><!-- End .banner-content --></div><!-- End .banner --></div><!-- End .col-lg-6 --></div><!-- End .row --><div class=\"row justify-content-center\"><div class=\"col-md-6 col-lg-4\"><div class=\"banner banner-overlay text-white\"><a href=\"#\"><img src=\"assets/images/demos/demo-7/banners/banner-9.jpg\" alt=\"Banner\"></a><div class=\"banner-content banner-content-right\"><h4 class=\"banner-subtitle\"><a href=\"#\">Necklace</a></h4><!-- End .banner-subtitle --><h3 class=\"banner-title\"><a href=\"#\">Summer<br>sale -70% off</a></h3><!-- End .banner-title --><a href=\"#\" class=\"btn underline btn-outline-white-3 banner-link\">Shop Now</a></div><!-- End .banner-content --></div><!-- End .banner --></div><!-- End .col-lg-4 --><div class=\"col-md-6 col-lg-4\"><div class=\"banner banner-overlay color-grey\"><a href=\"#\"><img src=\"assets/images/demos/demo-7/banners/banner-12.jpg\" alt=\"Banner\"></a><div class=\"banner-content\"><h4 class=\"banner-subtitle\"><a href=\"#\">Accessories</a></h4><!-- End .banner-subtitle --><h3 class=\"banner-title\"><a href=\"#\">2024 Winter<br>up to 50% off</a></h3><!-- End .banner-title --><a href=\"#\" class=\"btn underline banner-link\">Shop Now</a></div><!-- End .banner-content --></div><!-- End .banner --></div><!-- End .col-lg-4 --><div class=\"col-md-6 col-lg-4\"><div class=\"banner banner-overlay text-white\"><a href=\"#\"><img src=\"assets/images/demos/demo-7/banners/banner-11.jpg\" alt=\"Banner\"></a><div class=\"banner-content banner-content-right mr\"><h4 class=\"banner-subtitle\"><a href=\"#\">New in</a></h4><!-- End .banner-subtitle --><h3 class=\"banner-title\"><a href=\"#\">Kid’s<br>dress</a></h3><!-- End .banner-title --><a href=\"#\" class=\"btn underline btn-outline-white-3 banner-link\">Shop Now</a></div><!-- End .banner-content --></div><!-- End .banner --></div><!-- End .col-lg-4 --></div><!-- End .row --></div><!-- End .container-fluid --> <div class=\"bg-light-2 pt-6 pb-6 featured\"><div class=\"container-fluid\"><div class=\"heading heading-center mb-3\"><h2 class=\"title\">FEATURED PRODUCTS</h2><!-- End .title --><ul class=\"nav nav-pills justify-content-center\" role=\"tablist\"><li class=\"nav-item\"><a class=\"nav-link active\" id=\"featured-kid-link\" data-toggle=\"tab\" href=\"#featured-kid-tab\" role=\"tab\" aria-controls=\"featured-kid-tab\" aria-selected=\"true\">Kid’s Clothing</a></li><li class=\"nav-item\"><a class=\"nav-link\" id=\"featured-women-link\" data-toggle=\"tab\" href=\"#featured-women-tab\" role=\"tab\" aria-controls=\"featured-women-tab\" aria-selected=\"false\">Women's Accessories</a></li></ul></div><!-- End .heading --><div class=\"tab-content tab-content-carousel\"><div class=\"tab-pane p-0 fade show active\" id=\"featured-kid-tab\" role=\"tabpanel\" aria-labelledby=\"featured-kid-link\"><div class=\"owl-carousel owl-simple carousel-equal-height carousel-with-shadow\" data-toggle=\"owl\" data-owl-options=\"{\n                                &#34;nav&#34;: false, \n                                &#34;dots&#34;: true,\n                                &#34;margin&#34;: 20,\n                                &#34;loop&#34;: false,\n                                &#34;responsive&#34;: {\n                                    &#34;0&#34;: {\n                                        &#34;items&#34;:2\n                                    },\n                                    &#34;480&#34;: {\n                                        &#34;items&#34;:2\n                                    },\n                                    &#34;768&#34;: {\n                                        &#34;items&#34;:3\n                                    },\n                                    &#34;992&#34;: {\n                                        &#34;items&#34;:4\n                                    },\n                                    &#34;1200&#34;: {\n                                        &#34;items&#34;:5,\n                                        &#34;nav&#34;: true\n                                    }\n                                }\n                            }\"><div class=\"product product-7 text-center\"><figure class=\"product-media\"><span class=\"product-label label-circle label-sale\">Sale</span> <a href=\"#\"><img src=\"assets/images/products/kids/batch_1/Alia_Dress-4.jpeg\" alt=\"Product image\" class=\"product-image\"> <img src=\"assets/images/products/kids/batch_1/Alia_Dress-5.jpeg\" alt=\"Product image\" class=\"product-image-hover\"></a><div class=\"product-action-vertical\"><a href=\"#\" class=\"btn-product-icon btn-wishlist btn-expandable\"><span>add to wishlist</span></a> <a href=\"/popup/quickView.html\" class=\"btn-product-icon btn-quickview\" title=\"Quick view\"><span>Quick view</span></a></div><!-- End .product-action-vertical --><div class=\"product-action\"><a href=\"#\" class=\"btn-product btn-cart\"><span>add to cart</span></a></div><!-- End .product-action --></figure><!-- End .product-media --><div class=\"product-body\"><h3 class=\"product-title\"><a href=\"#\">Alia Dress - Green</a></h3><!-- End .product-title --><div class=\"product-price\"><span class=\"new-price\">Rp144.000</span> <span class=\"old-price\">Rp289.00</span></div><!-- End .product-price --><div class=\"ratings-container\"><div class=\"ratings\"><div class=\"ratings-val\" style=\"width: 100%;\"></div><!-- End .ratings-val --></div><!-- End .ratings --><span class=\"ratings-text\">( 6 Reviews )</span></div><!-- End .rating-container --></div><!-- End .product-body --></div><!-- End .product --><div class=\"product product-7 text-center\"><figure class=\"product-media\"><span class=\"product-label label-circle label-sale\">Sale</span> <a href=\"#\"><img src=\"assets/images/products/kids/batch_1/Alia_Dress-6.jpeg\" alt=\"Product image\" class=\"product-image\"> <img src=\"assets/images/products/kids/batch_1/Alia_Dress-7.jpeg\" alt=\"Product image\" class=\"product-image-hover\"></a><div class=\"product-action-vertical\"><a href=\"#\" class=\"btn-product-icon btn-wishlist btn-expandable\"><span>add to wishlist</span></a> <a href=\"popup/quickView.html\" class=\"btn-product-icon btn-quickview\" title=\"Quick view\"><span>Quick view</span></a></div><!-- End .product-action-vertical --><div class=\"product-action\"><a href=\"#\" class=\"btn-product btn-cart\"><span>add to cart</span></a></div><!-- End .product-action --></figure><!-- End .product-media --><div class=\"product-body\"><h3 class=\"product-title\"><a href=\"#\">Alia Dress - Dusty Pink</a></h3><!-- End .product-title --><div class=\"product-price\"><span class=\"new-price\">Rp144.000</span> <span class=\"old-price\">Rp289.00</span></div><!-- End .product-price --><div class=\"ratings-container\"><div class=\"ratings\"><div class=\"ratings-val\" style=\"width: 100%;\"></div><!-- End .ratings-val --></div><!-- End .ratings --><span class=\"ratings-text\">( 6 Reviews )</span></div><!-- End .rating-container --></div><!-- End .product-body --></div><!-- End .product --><div class=\"product product-7 text-center\"><figure class=\"product-media\"><span class=\"product-label label-circle label-top\">Top</span> <span class=\"product-label label-circle label-sale\">Sale</span> <a href=\"#\"><img src=\"assets/images/products/kids/batch_1/Myra_Dress-1.jpg\" alt=\"Product image\" class=\"product-image\"> <img src=\"assets/images/products/kids/batch_1/Myra_Dress-2.jpeg\" alt=\"Product image\" class=\"product-image-hover\"></a><div class=\"product-action-vertical\"><a href=\"#\" class=\"btn-product-icon btn-wishlist btn-expandable\"><span>add to wishlist</span></a> <a href=\"popup/quickView.html\" class=\"btn-product-icon btn-quickview\" title=\"Quick view\"><span>Quick view</span></a></div><!-- End .product-action-vertical --><div class=\"product-action\"><a href=\"#\" class=\"btn-product btn-cart\"><span>add to cart</span></a></div><!-- End .product-action --></figure><!-- End .product-media --><div class=\"product-body\"><h3 class=\"product-title\"><a href=\"#\">Myra Dress - Blue</a></h3><!-- End .product-title --><div class=\"product-price\"><span class=\"new-price\">Rp295.000</span> <span class=\"old-price\">Rp349.000</span></div><!-- End .product-price --><div class=\"ratings-container\"><div class=\"ratings\"><div class=\"ratings-val\" style=\"width: 100%;\"></div><!-- End .ratings-val --></div><!-- End .ratings --><span class=\"ratings-text\">( 24 Reviews )</span></div><!-- End .rating-container --></div><!-- End .product-body --></div><!-- End .product --><div class=\"product product-7 text-center\"><figure class=\"product-media\"><span class=\"product-label label-circle label-top\">Top</span> <span class=\"product-label label-circle label-sale\">Sale</span> <a href=\"#\"><img src=\"assets/images/products/kids/batch_1/Myra_Dress-5.jpg\" alt=\"Product image\" class=\"product-image\"> <img src=\"assets/images/products/kids/batch_1/Myra_Dress-6.jpeg\" alt=\"Product image\" class=\"product-image-hover\"></a><div class=\"product-action-vertical\"><a href=\"#\" class=\"btn-product-icon btn-wishlist btn-expandable\"><span>add to wishlist</span></a> <a href=\"popup/quickView.html\" class=\"btn-product-icon btn-quickview\" title=\"Quick view\"><span>Quick view</span></a></div><!-- End .product-action-vertical --><div class=\"product-action\"><a href=\"#\" class=\"btn-product btn-cart\"><span>add to cart</span></a></div><!-- End .product-action --></figure><!-- End .product-media --><div class=\"product-body\"><h3 class=\"product-title\"><a href=\"#\">Myra Dress - Dusty Pink</a></h3><!-- End .product-title --><div class=\"product-price\"><span class=\"new-price\">Rp295.000</span> <span class=\"old-price\">Rp349.000</span></div><!-- End .product-price --><div class=\"ratings-container\"><div class=\"ratings\"><div class=\"ratings-val\" style=\"width: 100%;\"></div><!-- End .ratings-val --></div><!-- End .ratings --><span class=\"ratings-text\">( 24 Reviews )</span></div><!-- End .rating-container --></div><!-- End .product-body --></div><!-- End .product --><div class=\"product product-7 text-center\"><figure class=\"product-media\"><span class=\"product-label label-circle label-top\">Top</span> <a href=\"#\"><img src=\"assets/images/products/kids/batch_1/Cyra_Raya_Collection-2.jpeg\" alt=\"Product image\" class=\"product-image\"> <img src=\"assets/images/products/kids/batch_1/Cyra_Raya_Collection-1.jpeg\" alt=\"Product image\" class=\"product-image-hover\"></a><div class=\"product-action-vertical\"><a href=\"#\" class=\"btn-product-icon btn-wishlist btn-expandable\"><span>add to wishlist</span></a> <a href=\"popup/quickView.html\" class=\"btn-product-icon btn-quickview\" title=\"Quick view\"><span>Quick view</span></a></div><!-- End .product-action-vertical --><div class=\"product-action\"><a href=\"#\" class=\"btn-product btn-cart\"><span>add to cart</span></a></div><!-- End .product-action --></figure><!-- End .product-media --><div class=\"product-body\"><h3 class=\"product-title\"><a href=\"#\">Cyra Raya Collection</a></h3><!-- End .product-title --><div class=\"product-price\">Rp330.000</div><!-- End .product-price --><div class=\"ratings-container\"><div class=\"ratings\"><div class=\"ratings-val\" style=\"width: 100%;\"></div><!-- End .ratings-val --></div><!-- End .ratings --><span class=\"ratings-text\">( 4 Reviews )</span></div><!-- End .rating-container --></div><!-- End .product-body --></div><!-- End .product --></div><!-- End .owl-carousel --></div><!-- .End .tab-pane --><div class=\"tab-pane p-0 fade\" id=\"featured-women-tab\" role=\"tabpanel\" aria-labelledby=\"featured-women-link\"><div class=\"owl-carousel owl-simple carousel-equal-height carousel-with-shadow\" data-toggle=\"owl\" data-owl-options=\"{\n                                &#34;nav&#34;: false, \n                                &#34;dots&#34;: true,\n                                &#34;margin&#34;: 20,\n                                &#34;loop&#34;: false,\n                                &#34;responsive&#34;: {\n                                    &#34;0&#34;: {\n                                        &#34;items&#34;:1\n                                    },\n                                    &#34;480&#34;: {\n                                        &#34;items&#34;:2\n                                    },\n                                    &#34;768&#34;: {\n                                        &#34;items&#34;:3\n                                    },\n                                    &#34;992&#34;: {\n                                        &#34;items&#34;:4\n                                    },\n                                    &#34;1200&#34;: {\n                                        &#34;items&#34;:5,\n                                        &#34;nav&#34;: true\n                                    }\n                                }\n                            }\"><div class=\"product product-7 text-center\"><figure class=\"product-media\"><a href=\"#\"><img src=\"assets/images/products/womens/aksesoris/kalung/Goca-1.jpg\" alt=\"Product image\" class=\"product-image\"> <img src=\"assets/images/products/womens/aksesoris/kalung/Goca-4.jpg\" alt=\"Product image\" class=\"product-image-hover\"></a><div class=\"product-action-vertical\"><a href=\"#\" class=\"btn-product-icon btn-wishlist btn-expandable\"><span>add to wishlist</span></a> <a href=\"popup/quickView.html\" class=\"btn-product-icon btn-quickview\" title=\"Quick view\"><span>Quick view</span></a></div><!-- End .product-action-vertical --><div class=\"product-action\"><a href=\"#\" class=\"btn-product btn-cart\"><span>add to cart</span></a></div><!-- End .product-action --></figure><!-- End .product-media --><div class=\"product-body\"><h3 class=\"product-title\"><a href=\"#\">Goca | Kalung</a></h3><!-- End .product-title --><div class=\"product-price\">Rp50.000</div><!-- End .product-price --><div class=\"ratings-container\"><div class=\"ratings\"><div class=\"ratings-val\" style=\"width: 100%;\"></div><!-- End .ratings-val --></div><!-- End .ratings --><span class=\"ratings-text\">( 3 Reviews )</span></div><!-- End .rating-container --></div><!-- End .product-body --></div><!-- End .product --><div class=\"product product-7 text-center\"><figure class=\"product-media\"><a href=\"#\"><img src=\"assets/images/products/womens/aksesoris/kalung/Elo-1.jpg\" alt=\"Product image\" class=\"product-image\"> <img src=\"assets/images/products/womens/aksesoris/kalung/Elo-2.jpg\" alt=\"Product image\" class=\"product-image-hover\"></a><div class=\"product-action-vertical\"><a href=\"#\" class=\"btn-product-icon btn-wishlist btn-expandable\"><span>add to wishlist</span></a> <a href=\"popup/quickView.html\" class=\"btn-product-icon btn-quickview\" title=\"Quick view\"><span>Quick view</span></a></div><!-- End .product-action-vertical --><div class=\"product-action\"><a href=\"#\" class=\"btn-product btn-cart\"><span>add to cart</span></a></div><!-- End .product-action --></figure><!-- End .product-media --><div class=\"product-body\"><h3 class=\"product-title\"><a href=\"#\">Elo | Kalung</a></h3><!-- End .product-title --><div class=\"product-price\">Rp65.000</div><!-- End .product-price --><div class=\"ratings-container\"><div class=\"ratings\"><div class=\"ratings-val\" style=\"width: 100%;\"></div><!-- End .ratings-val --></div><!-- End .ratings --><span class=\"ratings-text\">( 3 Reviews )</span></div><!-- End .rating-container --></div><!-- End .product-body --></div><!-- End .product --></div><!-- End .owl-carousel --></div><!-- .End .tab-pane --></div><!-- End .tab-content --></div><!-- End .container-fluid --></div><!-- End .bg-light-2 pt-4 pb-4 --> <div class=\"mb-6\"></div><!-- End .mb-6 --> <div class=\"container-fluid\"><div class=\"heading heading-center mb-3\"><ul class=\"nav nav-pills justify-content-center\" role=\"tablist\"><li class=\"nav-item\"><a class=\"nav-link active\" id=\"new-arrivals-link\" data-toggle=\"tab\" href=\"#new-arrivals-tab\" role=\"tab\" aria-controls=\"new-arrivals-tab\" aria-selected=\"true\">New Arrivals</a></li><li class=\"nav-item\"><a class=\"nav-link\" id=\"last-stock-link\" data-toggle=\"tab\" href=\"#last-stock-tab\" role=\"tab\" aria-controls=\"last-stock-tab\" aria-selected=\"false\">Last Stock</a></li></ul></div><!-- End .heading --><div class=\"tab-content\"><div class=\"tab-pane p-0 fade show active\" id=\"new-arrivals-tab\" role=\"tabpanel\" aria-labelledby=\"new-arrivals-link\"><div class=\"products\"><div class=\"row justify-content-center\"><div class=\"col-6 col-md-4 col-lg-3 col-xl-5col\"><div class=\"product product-7 text-center\"><figure class=\"product-media\"><span class=\"product-label label-circle label-sale\">Sale</span> <a href=\"#\"><img src=\"assets/images/products/kids/batch_1/Alia_Dress-4.jpeg\" alt=\"Product image\" class=\"product-image\"> <img src=\"assets/images/products/kids/batch_1/Alia_Dress-5.jpeg\" alt=\"Product image\" class=\"product-image-hover\"></a><div class=\"product-action-vertical\"><a href=\"#\" class=\"btn-product-icon btn-wishlist btn-expandable\"><span>add to wishlist</span></a> <a href=\"/popup/quickView.html\" class=\"btn-product-icon btn-quickview\" title=\"Quick view\"><span>Quick view</span></a></div><!-- End .product-action-vertical --><div class=\"product-action\"><a href=\"#\" class=\"btn-product btn-cart\"><span>add to cart</span></a></div><!-- End .product-action --></figure><!-- End .product-media --><div class=\"product-body\"><h3 class=\"product-title\"><a href=\"#\">Alia Dress - Green</a></h3><!-- End .product-title --><div class=\"product-price\"><span class=\"new-price\">Rp144.000</span> <span class=\"old-price\">Rp289.00</span></div><!-- End .product-price --><div class=\"ratings-container\"><div class=\"ratings\"><div class=\"ratings-val\" style=\"width: 100%;\"></div><!-- End .ratings-val --></div><!-- End .ratings --><span class=\"ratings-text\">( 6 Reviews )</span></div><!-- End .rating-container --></div><!-- End .product-body --></div><!-- End .product --></div><div class=\"col-6 col-md-4 col-lg-3 col-xl-5col\"><div class=\"product product-7 text-center\"><figure class=\"product-media\"><span class=\"product-label label-circle label-sale\">Sale</span> <a href=\"#\"><img src=\"assets/images/products/kids/batch_1/Alia_Dress-6.jpeg\" alt=\"Product image\" class=\"product-image\"> <img src=\"assets/images/products/kids/batch_1/Alia_Dress-7.jpeg\" alt=\"Product image\" class=\"product-image-hover\"></a><div class=\"product-action-vertical\"><a href=\"#\" class=\"btn-product-icon btn-wishlist btn-expandable\"><span>add to wishlist</span></a> <a href=\"popup/quickView.html\" class=\"btn-product-icon btn-quickview\" title=\"Quick view\"><span>Quick view</span></a></div><!-- End .product-action-vertical --><div class=\"product-action\"><a href=\"#\" class=\"btn-product btn-cart\"><span>add to cart</span></a></div><!-- End .product-action --></figure><!-- End .product-media --><div class=\"product-body\"><h3 class=\"product-title\"><a href=\"#\">Alia Dress - Dusty Pink</a></h3><!-- End .product-title --><div class=\"product-price\"><span class=\"new-price\">Rp144.000</span> <span class=\"old-price\">Rp289.00</span></div><!-- End .product-price --><div class=\"ratings-container\"><div class=\"ratings\"><div class=\"ratings-val\" style=\"width: 100%;\"></div><!-- End .ratings-val --></div><!-- End .ratings --><span class=\"ratings-text\">( 6 Reviews )</span></div><!-- End .rating-container --></div><!-- End .product-body --></div><!-- End .product --></div></div><!-- End .row --></div><!-- End .products --></div><!-- .End .tab-pane --><div class=\"tab-pane p-0 fade\" id=\"last-stock-tab\" role=\"tabpanel\" aria-labelledby=\"last-stock-link\"><div class=\"products\"><div class=\"row justify-content-center\"><div class=\"col-6 col-md-4 col-lg-3 col-xl-5col\"><div class=\"product product-7 text-center\"><figure class=\"product-media\"><span class=\"product-label label-circle label-sale\">Sale</span> <a href=\"#\"><img src=\"assets/images/products/kids/batch_1/Myra_Dress-1.jpg\" alt=\"Product image\" class=\"product-image\"> <img src=\"assets/images/products/kids/batch_1/Myra_Dress-2.jpeg\" alt=\"Product image\" class=\"product-image-hover\"></a><div class=\"product-action-vertical\"><a href=\"#\" class=\"btn-product-icon btn-wishlist btn-expandable\"><span>add to wishlist</span></a> <a href=\"popup/quickView.html\" class=\"btn-product-icon btn-quickview\" title=\"Quick view\"><span>Quick view</span></a></div><!-- End .product-action-vertical --><div class=\"product-action\"><a href=\"#\" class=\"btn-product btn-cart\"><span>add to cart</span></a></div><!-- End .product-action --></figure><!-- End .product-media --><div class=\"product-body\"><h3 class=\"product-title\"><a href=\"#\">Myra Dress - Blue</a></h3><!-- End .product-title --><div class=\"product-price\"><span class=\"new-price\">Rp295.000</span> <span class=\"old-price\">Rp349.000</span></div><!-- End .product-price --><div class=\"ratings-container\"><div class=\"ratings\"><div class=\"ratings-val\" style=\"width: 100%;\"></div><!-- End .ratings-val --></div><!-- End .ratings --><span class=\"ratings-text\">( 24 Reviews )</span></div><!-- End .rating-container --></div><!-- End .product-body --></div><!-- End .product --></div><div class=\"col-6 col-md-4 col-lg-3 col-xl-5col\"><div class=\"product product-7 text-center\"><figure class=\"product-media\"><span class=\"product-label label-circle label-sale\">Sale</span> <a href=\"#\"><img src=\"assets/images/products/kids/batch_1/Myra_Dress-5.jpg\" alt=\"Product image\" class=\"product-image\"> <img src=\"assets/images/products/kids/batch_1/Myra_Dress-6.jpeg\" alt=\"Product image\" class=\"product-image-hover\"></a><div class=\"product-action-vertical\"><a href=\"#\" class=\"btn-product-icon btn-wishlist btn-expandable\"><span>add to wishlist</span></a> <a href=\"popup/quickView.html\" class=\"btn-product-icon btn-quickview\" title=\"Quick view\"><span>Quick view</span></a></div><!-- End .product-action-vertical --><div class=\"product-action\"><a href=\"#\" class=\"btn-product btn-cart\"><span>add to cart</span></a></div><!-- End .product-action --></figure><!-- End .product-media --><div class=\"product-body\"><h3 class=\"product-title\"><a href=\"#\">Myra Dress - Dusty Pink</a></h3><!-- End .product-title --><div class=\"product-price\"><span class=\"new-price\">Rp295.000</span> <span class=\"old-price\">Rp349.000</span></div><!-- End .product-price --><div class=\"ratings-container\"><div class=\"ratings\"><div class=\"ratings-val\" style=\"width: 100%;\"></div><!-- End .ratings-val --></div><!-- End .ratings --><span class=\"ratings-text\">( 24 Reviews )</span></div><!-- End .rating-container --></div><!-- End .product-body --></div><!-- End .product --></div><!-- End .col-sm-6 col-md-4 col-lg-3 --></div><!-- End .row --></div><!-- End .products --></div><!-- .End .tab-pane --></div><!-- End .tab-content --><div class=\"more-container text-center mt-2\"><a href=\"#\" class=\"btn btn-outline-dark-3 btn-more\"><span>Load more</span><i class=\"icon-long-arrow-right\"></i></a></div><!-- End .more-container --></div><!-- End .container-fluid --> <div class=\"bg-light-2 pt-7 pb-6 testimonials\"><div class=\"container\"><h2 class=\"title text-center mb-2\">Our Customers Say</h2><!-- End .title text-center --><div class=\"owl-carousel owl-simple owl-testimonials\" data-toggle=\"owl\" data-owl-options=\"{\n                        &#34;nav&#34;: false, \n                        &#34;dots&#34;: true,\n                        &#34;margin&#34;: 20,\n                        &#34;loop&#34;: false,\n                        &#34;responsive&#34;: {\n                            &#34;1200&#34;: {\n                                &#34;nav&#34;: true\n                            }\n                        }\n                    }\"><blockquote class=\"testimonial testimonial-icon text-center\"><p class=\"lead\">“Alia Dress - Dusty Pink”</p><!-- End .lead --><p>“ Love at first sight sama dress2 di Fralla, sampai galau mau checkout yang mana karna cantik2 semua. Bahan, model dan warna super cantik. Size dan panjangnya pas banget di anakku (usia 2th 3bln, BB 12kg-an). Bisa buat occasion apa aja. Terima kasih free ikat rambut cantiknya 🩷 ”</p><cite>farahditalrst, <span>Customer</span></cite></blockquote><!-- End .testimonial --><blockquote class=\"testimonial testimonial-icon text-center\"><p class=\"lead\">“Sofia Dress”</p><!-- End .lead --><p>“ Bajunya bagus bgt modelnyaa lucu unik.. knp baru nemu toko ini skrg sih.. luv bgt bajunya gk pasaran🫶🏻🫶🏻”</p><cite>syafiyah, <span>Customer</span></cite></blockquote><!-- End .testimonial --><blockquote class=\"testimonial testimonial-icon text-center\"><p class=\"lead\">“Kaftan Anak”</p><!-- End .lead --><p>“ Cantikkk banget , bahannyaa juga bagus banget , puaaaass 😍 ”</p><cite>indaharprtw, <span>Customer</span></cite></blockquote><!-- End .testimonial --><blockquote class=\"testimonial testimonial-icon text-center\"><p class=\"lead\">“Cyra | Raya Collection”</p><!-- End .lead --><p>“ Ngga pernah gagal koleksi dr fralla bahannya good quality, warnanya cantikkk cantikkkk modelnya mewahhh 💞💞 thanks seller pertahankan model dan bahannya untuk seri selanjutnya badabesttt 💞 ”</p><cite>febyhaerani, <span>Customer</span></cite></blockquote><!-- End .testimonial --><blockquote class=\"testimonial testimonial-icon text-center\"><p class=\"lead\">Zayra Cream | Kaftan”</p><!-- End .lead --><p>“ Bagusss bgtttttt masya allah... nunggu yg warna grey semoga dibikin ya . Trims ”</p><cite>miadacosta, <span>Customer</span></cite></blockquote><!-- End .testimonial --><blockquote class=\"testimonial testimonial-icon text-center\"><p class=\"lead\">Bandi | Sikka</p><!-- End .lead --><p>“ Packingnya aman dan rapi banget. Aman karna pakai bubble wrap, admin responsif, barang sesuai deskripsi, dikirim sesuai pilihan. Warnanya mewah gitu karna dari bahan silk. Puassss bgt ”</p><cite>dinindi, <span>Customer</span></cite></blockquote><!-- End .testimonial --><blockquote class=\"testimonial testimonial-icon text-center\"><p class=\"lead\">Keyla Dress</p><!-- End .lead --><p>“ udah punya beberapa koleksi fralla <br>MasyaAllah bagus2 semua <br>anakku jadi tambah cantik pake dressnya<br>trimakasih Fralla kids  semoga makin sukses 😚😉 ”</p><cite>nnoovvaa_arsy, <span>Customer</span></cite></blockquote><!-- End .testimonial --></div><!-- End .testimonials-slider owl-carousel --></div><!-- End .container --></div><!-- End .bg-light pt-5 pb-5 --> <div class=\"brands-border owl-carousel owl-simple\" data-toggle=\"owl\" data-owl-options=\"{\n                &#34;nav&#34;: false, \n                &#34;dots&#34;: false,\n                &#34;margin&#34;: 0,\n                &#34;loop&#34;: false,\n                &#34;responsive&#34;: {\n                    &#34;0&#34;: {\n                        &#34;items&#34;:2\n                    },\n                    &#34;420&#34;: {\n                        &#34;items&#34;:3\n                    },\n                    &#34;600&#34;: {\n                        &#34;items&#34;:4\n                    },\n                    &#34;900&#34;: {\n                        &#34;items&#34;:5\n                    },\n                    &#34;1024&#34;: {\n                        &#34;items&#34;:6\n                    },\n                    &#34;1360&#34;: {\n                        &#34;items&#34;:7\n                    }\n                }\n            }\"><a href=\"#\" class=\"brand\"><img src=\"assets/images/brands/1.png\" alt=\"Fralla Official\"></a> <a href=\"#\" class=\"brand\"><img src=\"assets/images/brands/2.png\" alt=\"Fralla Kids\"></a> <a href=\"#\" class=\"brand\"><img src=\"assets/images/brands/3.png\" alt=\"Brand Name\"></a> <a href=\"#\" class=\"brand\"></a> <a href=\"#\" class=\"brand\"></a> <a href=\"#\" class=\"brand\"></a> <a href=\"#\" class=\"brand\"></a></div><!-- End .owl-carousel -->")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = components.Layout("", token).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
