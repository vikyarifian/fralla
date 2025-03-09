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

func About(token dto.Token) templ.Component {
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
			templ_7745c5c3_Err = components.Menu(token).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "          <div class=\"mb-6 mb-lg-8\"><div class=\"container\"><div class=\"row\"><div class=\"col-lg-8 mb-3 mb-lg-0 pr-10 pl-10\" style=\"padding: 5rem 5rem 0 5rem;font-size:1.4rem;color:#000;\"><p class=\"mb-2\"><b>Fralla</b> adalah destinasi yang menghadirkan koleksi baju anak berkualitas dan aksesoris wanita elegan dalam satu tempat. Kami memahami bahwa kenyamanan dan gaya adalah prioritas, baik untuk si kecil maupun para wanita yang ingin tampil modis dalam setiap kesempatan.<br><br>Kenapa Belanja di Fralla?<ul style=\"list-style-type: circle;\"><li>Baju Anak Nyaman & Stylish – Koleksi pakaian anak dengan desain menggemaskan, bahan lembut, dan aman untuk kulit si kecil.</li><li>Aksesoris Wanita yang Elegan – Dari kalung, gelang dan aksesoris lainnya untuk melengkapi tampilan Anda.</li><li>Kualitas Terbaik – Produk dipilih dengan standar tinggi untuk memastikan kepuasan pelanggan.</li><li>Belanja Praktis & Aman – Proses belanja mudah dan pengiriman cepat ke seluruh Indonesia.</li></ul><br>Fralla – Untuk Si Kecil & Untuk Anda, dengan Gaya dan Cinta! </p></div><!-- End .col-lg-5 --><div class=\"col-lg-6 offset-lg-1\"><div class=\"about-images\"></div><!-- End .about-images --></div><!-- End .col-lg-6 --></div><!-- End .row --></div><!-- End .container --></div><!-- End .bg-light-2 pt-6 pb-6 -->")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = components.Layout("About", token).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
