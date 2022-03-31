package entities

type Model_setting struct {
	StartMaintenance string `json:"maintenance_start"`
	EndMaintenance   string `json:"maintenance_end"`
}
type Model_domain struct {
	Domain string `json:"domain"`
}
type Model_mclientpasaran struct {
	PasaranId             string `json:"pasaran_id"`
	PasaranTogel          string `json:"pasaran_togel"`
	PasaranPeriode        string `json:"pasaran_periode"`
	PasaranTglKeluaran    string `json:"pasaran_tglkeluaran"`
	Pasaranmarketclose    string `json:"pasaran_marketclose"`
	Pasaranmarketschedule string `json:"pasaran_marketschedule"`
	Pasaranmarketopen     string `json:"pasaran_marketopen"`
	Pasaranjamtutup       string `json:"pasaran_jamtutup"`
	Pasaranjamopen        string `json:"pasaran_jamopen"`
	Pasarannote           string `json:"pasaran_note"`
	Pasaranurl            string `json:"pasaran_url"`
	Pasaranhari           string `json:"pasaran_hari"`
	PasaranStatus         string `json:"pasaran_status"`
}
type Model_mclientpasaranResult struct {
	No      int    `json:"no"`
	Date    string `json:"date"`
	Periode string `json:"periode"`
	Result  string `json:"result"`
}
type Model_mclientpasaranResultAll struct {
	No           int    `json:"no"`
	Date         string `json:"date"`
	Periode      string `json:"periode"`
	Pasaran      string `json:"pasaran"`
	Pasaran_code string `json:"pasaran_code"`
	Result       string `json:"result"`
}
type Model_mclientpasaranCheckPasaran struct {
	PasaranIdtansaction string `json:"pasaran_idtransaction"`
	PasaranName         string `json:"pasaran_name"`
	PasaranPeriode      string `json:"pasaran_periode"`
	PasaranIdcomp       string `json:"pasaran_idcomp"`
	PasaranStatus       string `json:"pasaran_status"`
}
type Model_mpasarantogel432 struct {
	Min_bet            float32 `json:"min_bet"`
	Max4d_bet          float32 `json:"max4d_bet"`
	Max3d_bet          float32 `json:"max3d_bet"`
	Max3dd_bet         float32 `json:"max3dd_bet"`
	Max2d_bet          float32 `json:"max2d_bet"`
	Max2dd_bet         float32 `json:"max2dd_bet"`
	Max2dt_bet         float32 `json:"max2dt_bet"`
	Disc4d_bet         float32 `json:"disc4d_bet"`
	Disc3d_bet         float32 `json:"disc3d_bet"`
	Disc3dd_bet        float32 `json:"disc3dd_bet"`
	Disc2d_bet         float32 `json:"disc2d_bet"`
	Disc2dd_bet        float32 `json:"disc2dd_bet"`
	Disc2dt_bet        float32 `json:"disc2dt_bet"`
	Win4d_bet          float32 `json:"win4d_bet"`
	Win3d_bet          float32 `json:"win3d_bet"`
	Win3dd_bet         float32 `json:"win3dd_bet"`
	Win2d_bet          float32 `json:"win2d_bet"`
	Win2dd_bet         float32 `json:"win2dd_bet"`
	Win2dt_bet         float32 `json:"win2dt_bet"`
	Win4dnodiskon_bet  float32 `json:"win4dnodiskon_bet"`
	Win3dnodiskon_bet  float32 `json:"win3dnodiskon_bet"`
	Win3ddnodiskon_bet float32 `json:"win3ddnodiskon_bet"`
	Win2dnodiskon_bet  float32 `json:"win2dnodiskon_bet"`
	Win2ddnodiskon_bet float32 `json:"win2ddnodiskon_bet"`
	Win2dtnodiskon_bet float32 `json:"win2dtnodiskon_bet"`
	Win4dbb_kena_bet   float32 `json:"win4dbb_kena_bet"`
	Win3dbb_kena_bet   float32 `json:"win3dbb_kena_bet"`
	Win3ddbb_kena_bet  float32 `json:"win3ddbb_kena_bet"`
	Win2dbb_kena_bet   float32 `json:"win2dbb_kena_bet"`
	Win2ddbb_kena_bet  float32 `json:"win2ddbb_kena_bet"`
	Win2dtbb_kena_bet  float32 `json:"win2dtbb_kena_bet"`
	Win4dbb_bet        float32 `json:"win4dbb_bet"`
	Win3dbb_bet        float32 `json:"win3dbb_bet"`
	Win3ddbb_bet       float32 `json:"win3ddbb_bet"`
	Win2dbb_bet        float32 `json:"win2dbb_bet"`
	Win2ddbb_bet       float32 `json:"win2ddbb_bet"`
	Win2dtbb_bet       float32 `json:"win2dtbb_bet"`
	Limittotal4d_bet   float32 `json:"limittotal4d_bet"`
	Limittotal3d_bet   float32 `json:"limittotal3d_bet"`
	Limittotal3dd_bet  float32 `json:"limittotal3dd_bet"`
	Limittotal2d_bet   float32 `json:"limittotal2d_bet"`
	Limittotal2dd_bet  float32 `json:"limittotal2dd_bet"`
	Limittotal2dt_bet  float32 `json:"limittotal2dt_bet"`
	Limitline_4d       uint32  `json:"limitline_4d"`
	Limitline_3d       uint32  `json:"limitline_3d"`
	Limitline_3dd      uint32  `json:"limitline_3dd"`
	Limitline_2d       uint32  `json:"limitline_2d"`
	Limitline_2dd      uint32  `json:"limitline_2dd"`
	Limitline_2dt      uint32  `json:"limitline_2dt"`
	Bbfs               uint8   `json:"bbfs"`
}
type Model_mpasarantogelColok struct {
	Min_bet_colokbebas        float32 `json:"min_bet_colokbebas"`
	Max_bet_colokbebas        float32 `json:"max_bet_colokbebas"`
	Disc_bet_colokbebas       float32 `json:"disc_bet_colokbebas"`
	Win_bet_colokbebas        float32 `json:"win_bet_colokbebas"`
	Limittotal_bet_colokbebas float32 `json:"limittotal_bet_colokbebas"`
	Min_bet_colokmacau        float32 `json:"min_bet_colokmacau"`
	Max_bet_colokmacau        float32 `json:"max_bet_colokmacau"`
	Disc_bet_colokmacau       float32 `json:"disc_bet_colokmacau"`
	Win_bet_colokmacau        float32 `json:"win_bet_colokmacau"`
	Win3_bet_colokmacau       float32 `json:"win3_bet_colokmacau"`
	Win4_bet_colokmacau       float32 `json:"win4_bet_colokmacau"`
	Limittotal_bet_colokmacau float32 `json:"limittotal_bet_colokmacau"`
	Min_bet_coloknaga         float32 `json:"min_bet_coloknaga"`
	Max_bet_coloknaga         float32 `json:"max_bet_coloknaga"`
	Disc_bet_coloknaga        float32 `json:"disc_bet_coloknaga"`
	Win_bet_coloknaga         float32 `json:"win_bet_coloknaga"`
	Win4_bet_coloknaga        float32 `json:"win4_bet_coloknaga"`
	Limittotal_bet_coloknaga  float32 `json:"limittotal_bet_coloknaga"`
	Min_bet_colokjitu         float32 `json:"min_bet_colokjitu"`
	Max_bet_colokjitu         float32 `json:"max_bet_colokjitu"`
	Disc_bet_colokjitu        float32 `json:"disc_bet_colokjitu"`
	Winas_bet_colokjitu       float32 `json:"winas_bet_colokjitu"`
	Winkop_bet_colokjitu      float32 `json:"winkop_bet_colokjitu"`
	Winkepala_bet_colokjitu   float32 `json:"winkepala_bet_colokjitu"`
	Winekor_bet_colokjitu     float32 `json:"winekor_bet_colokjitu"`
	Limittotal_bet_colokjitu  float32 `json:"limittotal_bet_colokjitu"`
}
type Model_pasarantogel5050 struct {
	Min_bet_5050umum                       float32 `json:"min_bet_5050umum"`
	Max_bet_5050umum                       float32 `json:"max_bet_5050umum"`
	Keibesar_bet_5050umum                  float32 `json:"keibesar_bet_5050umum"`
	Keikecil_bet_5050umum                  float32 `json:"keikecil_bet_5050umum"`
	Keigenap_bet_5050umum                  float32 `json:"keigenap_bet_5050umum"`
	Keiganjil_bet_5050umum                 float32 `json:"keiganjil_bet_5050umum"`
	Keitengah_bet_5050umum                 float32 `json:"keitengah_bet_5050umum"`
	Keitepi_bet_5050umum                   float32 `json:"keitepi_bet_5050umum"`
	Discbesar_bet_5050umum                 float32 `json:"discbesar_bet_5050umum"`
	Disckecil_bet_5050umum                 float32 `json:"disckecil_bet_5050umum"`
	Discgenap_bet_5050umum                 float32 `json:"discgenap_bet_5050umum"`
	Discganjil_bet_5050umum                float32 `json:"discganjil_bet_5050umum"`
	Disctengah_bet_5050umum                float32 `json:"disctengah_bet_5050umum"`
	Disctepi_bet_5050umum                  float32 `json:"disctepi_bet_5050umum"`
	Limittotal_bet_5050umum                float32 `json:"limittotal_bet_5050umum"`
	Min_bet_5050special                    float32 `json:"min_bet_5050special"`
	Max_bet_5050special                    float32 `json:"max_bet_5050special"`
	Keiasganjil_bet_5050special            float32 `json:"keiasganjil_bet_5050special"`
	Keiasgenap_bet_5050special             float32 `json:"keiasgenap_bet_5050special"`
	Keiasbesar_bet_5050special             float32 `json:"keiasbesar_bet_5050special"`
	Keiaskecil_bet_5050special             float32 `json:"keiaskecil_bet_5050special"`
	Keikopganjil_bet_5050special           float32 `json:"keikopganjil_bet_5050special"`
	Keikopgenap_bet_5050special            float32 `json:"keikopgenap_bet_5050special"`
	Keikopbesar_bet_5050special            float32 `json:"keikopbesar_bet_5050special"`
	Keikopkecil_bet_5050special            float32 `json:"keikopkecil_bet_5050special"`
	Keikepalaganjil_bet_5050special        float32 `json:"keikepalaganjil_bet_5050special"`
	Keikepalagenap_bet_5050special         float32 `json:"keikepalagenap_bet_5050special"`
	Keikepalabesar_bet_5050special         float32 `json:"keikepalabesar_bet_5050special"`
	Keikepalakecil_bet_5050special         float32 `json:"keikepalakecil_bet_5050special"`
	Keiekorganjil_bet_5050special          float32 `json:"keiekorganjil_bet_5050special"`
	Keiekorgenap_bet_5050special           float32 `json:"keiekorgenap_bet_5050special"`
	Keiekorbesar_bet_5050special           float32 `json:"keiekorbesar_bet_5050special"`
	Keiekorkecil_bet_5050special           float32 `json:"keiekorkecil_bet_5050special"`
	Discasganjil_bet_5050special           float32 `json:"discasganjil_bet_5050special"`
	Discasgenap_bet_5050special            float32 `json:"discasgenap_bet_5050special"`
	Discasbesar_bet_5050special            float32 `json:"discasbesar_bet_5050special"`
	Discaskecil_bet_5050special            float32 `json:"discaskecil_bet_5050special"`
	Disckopganjil_bet_5050special          float32 `json:"disckopganjil_bet_5050special"`
	Disckopgenap_bet_5050special           float32 `json:"disckopgenap_bet_5050special"`
	Disckopbesar_bet_5050special           float32 `json:"disckopbesar_bet_5050special"`
	Disckopkecil_bet_5050special           float32 `json:"disckopkecil_bet_5050special"`
	Disckepalaganjil_bet_5050special       float32 `json:"disckepalaganjil_bet_5050special"`
	Disckepalagenap_bet_5050special        float32 `json:"disckepalagenap_bet_5050special"`
	Disckepalabesar_bet_5050special        float32 `json:"disckepalabesar_bet_5050special"`
	Disckepalakecil_bet_5050special        float32 `json:"disckepalakecil_bet_5050special"`
	Discekorganjil_bet_5050special         float32 `json:"discekorganjil_bet_5050special"`
	Discekorgenap_bet_5050special          float32 `json:"discekorgenap_bet_5050special"`
	Discekorbesar_bet_5050special          float32 `json:"discekorbesar_bet_5050special"`
	Discekorkecil_bet_5050special          float32 `json:"discekorkecil_bet_5050special"`
	Limittotal_bet_5050special             float32 `json:"limittotal_bet_5050special"`
	Min_bet_5050kombinasi                  float32 `json:"min_bet_5050kombinasi"`
	Max_bet_5050kombinasi                  float32 `json:"max_bet_5050kombinasi"`
	Kei_belakangmono_bet_5050kombinasi     float32 `json:"kei_belakangmono_bet_5050kombinasi"`
	Kei_belakangstereo_bet_5050kombinasi   float32 `json:"kei_belakangstereo_bet_5050kombinasi"`
	Kei_belakangkembang_bet_5050kombinasi  float32 `json:"kei_belakangkembang_bet_5050kombinasi"`
	Kei_belakangkempis_bet_5050kombinasi   float32 `json:"kei_belakangkempis_bet_5050kombinasi"`
	Kei_belakangkembar_bet_5050kombinasi   float32 `json:"kei_belakangkembar_bet_5050kombinasi"`
	Kei_tengahmono_bet_5050kombinasi       float32 `json:"kei_tengahmono_bet_5050kombinasi"`
	Kei_tengahstereo_bet_5050kombinasi     float32 `json:"kei_tengahstereo_bet_5050kombinasi"`
	Kei_tengahkembang_bet_5050kombinasi    float32 `json:"kei_tengahkembang_bet_5050kombinasi"`
	Kei_tengahkempis_bet_5050kombinasi     float32 `json:"kei_tengahkempis_bet_5050kombinasi"`
	Kei_tengahkembar_bet_5050kombinasi     float32 `json:"kei_tengahkembar_bet_5050kombinasi"`
	Kei_depanmono_bet_5050kombinasi        float32 `json:"kei_depanmono_bet_5050kombinasi"`
	Kei_depanstereo_bet_5050kombinasi      float32 `json:"kei_depanstereo_bet_5050kombinasi"`
	Kei_depankembang_bet_5050kombinasi     float32 `json:"kei_depankembang_bet_5050kombinasi"`
	Kei_depankempis_bet_5050kombinasi      float32 `json:"kei_depankempis_bet_5050kombinasi"`
	Kei_depankembar_bet_5050kombinasi      float32 `json:"kei_depankembar_bet_5050kombinasi"`
	Disc_belakangmono_bet_5050kombinasi    float32 `json:"disc_belakangmono_bet_5050kombinasi"`
	Disc_belakangstereo_bet_5050kombinasi  float32 `json:"disc_belakangstereo_bet_5050kombinasi"`
	Disc_belakangkembang_bet_5050kombinasi float32 `json:"disc_belakangkembang_bet_5050kombinasi"`
	Disc_belakangkempis_bet_5050kombinasi  float32 `json:"disc_belakangkempis_bet_5050kombinasi"`
	Disc_belakangkembar_bet_5050kombinasi  float32 `json:"disc_belakangkembar_bet_5050kombinasi"`
	Disc_tengahmono_bet_5050kombinasi      float32 `json:"disc_tengahmono_bet_5050kombinasi"`
	Disc_tengahstereo_bet_5050kombinasi    float32 `json:"disc_tengahstereo_bet_5050kombinasi"`
	Disc_tengahkembang_bet_5050kombinasi   float32 `json:"disc_tengahkembang_bet_5050kombinasi"`
	Disc_tengahkempis_bet_5050kombinasi    float32 `json:"disc_tengahkempis_bet_5050kombinasi"`
	Disc_tengahkembar_bet_5050kombinasi    float32 `json:"disc_tengahkembar_bet_5050kombinasi"`
	Disc_depanmono_bet_5050kombinasi       float32 `json:"disc_depanmono_bet_5050kombinasi"`
	Disc_depanstereo_bet_5050kombinasi     float32 `json:"disc_depanstereo_bet_5050kombinasi"`
	Disc_depankembang_bet_5050kombinasi    float32 `json:"disc_depankembang_bet_5050kombinasi"`
	Disc_depankempis_bet_5050kombinasi     float32 `json:"disc_depankempis_bet_5050kombinasi"`
	Disc_depankembar_bet_5050kombinasi     float32 `json:"disc_depankembar_bet_5050kombinasi"`
	Limittotal_bet_5050kombinasi           float32 `json:"limittotal_bet_5050kombinasi"`
}
type Model_pasarantogelMacauKombinasi struct {
	Min_bet     float32 `json:"min_bet"`
	Max_bet     float32 `json:"max_bet"`
	Win_bet     float32 `json:"win_bet"`
	Diskon_bet  float32 `json:"diskon_bet"`
	Limit_total float32 `json:"limit_total"`
}
type Model_mpasarantogelDasar struct {
	Min_bet         float32 `json:"min_bet"`
	Max_bet         float32 `json:"max_bet"`
	Kei_besar_bet   float32 `json:"kei_besar_bet"`
	Kei_kecil_bet   float32 `json:"kei_kecil_bet"`
	Kei_genap_bet   float32 `json:"kei_genap_bet"`
	Kei_ganjil_bet  float32 `json:"kei_ganjil_bet"`
	Disc_besar_bet  float32 `json:"disc_besar_bet"`
	Disc_kecil_bet  float32 `json:"disc_kecil_bet"`
	Disc_genap_bet  float32 `json:"disc_genap_bet"`
	Disc_ganjil_bet float32 `json:"disc_ganjil_bet"`
	Limit_total     float32 `json:"limit_total"`
}
type Model_mpasarantogelShio struct {
	Min_bet     float32 `json:"min_bet"`
	Max_bet     float32 `json:"max_bet"`
	Win_bet     float32 `json:"win_bet"`
	Diskon_bet  float32 `json:"diskon_bet"`
	Limit_total float32 `json:"limit_total"`
}
type Model_mpasaranLimit struct {
	Total_4d  int `json:"total_4d"`
	Total_3d  int `json:"total_3d"`
	Total_3dd int `json:"total_3dd"`
	Total_2d  int `json:"total_2d"`
	Total_2dd int `json:"total_2dd"`
	Total_2dt int `json:"total_2dt"`
}
type Model_mlistinvoicebet struct {
	Tanggal   string  `json:"tanggal"`
	Tipe      string  `json:"tipe"`
	Permainan string  `json:"permainan"`
	Periode   string  `json:"periode"`
	Nomor     string  `json:"nomor"`
	Bet       int     `json:"bet"`
	Diskon    float32 `json:"diskon"`
	Kei       float32 `json:"kei"`
	Bayar     int     `json:"bayar"`
	Win       int     `json:"win"`
	Menang    int     `json:"menang"`
}
type Model_mgroupinvoicebetPermainan struct {
	Permainan string `json:"permainan"`
}
type Model_mlistinvoicebetid struct {
	No        int     `json:"no"`
	Status    string  `json:"status"`
	Tipe      string  `json:"tipe"`
	Permainan string  `json:"permainan"`
	Nomor     string  `json:"nomor"`
	Bet       int     `json:"bet"`
	Diskon    float32 `json:"diskon"`
	Kei       float32 `json:"kei"`
	Bayar     int     `json:"bayar"`
	Win       int     `json:"win"`
}
type Model_mlistsipperiode struct {
	Tanggal         string `json:"tglkeluaran"`
	Idinvoice       string `json:"idinvoice"`
	Periode         string `json:"periode"`
	Totalbet        int    `json:"totalbet"`
	Totalbayar      int    `json:"totalbayar"`
	Totalwin        int    `json:"totalwin"`
	Totallose       int    `json:"totallose"`
	Status          string `json:"status"`
	Color_lost      string `json:"color_lost"`
	Background      string `json:"background"`
	Color_totallose string `json:"color_totallose"`
}
type Model_mlistsipperiodeall struct {
	Tanggal         string `json:"tglkeluaran"`
	Idinvoice       string `json:"idinvoice"`
	Pasaran         string `json:"pasaran"`
	Periode         string `json:"periode"`
	Totalbet        int    `json:"totalbet"`
	Totalbayar      int    `json:"totalbayar"`
	Totalwin        int    `json:"totalwin"`
	Totallose       int    `json:"totallose"`
	Status          string `json:"status"`
	Color_lost      string `json:"color_lost"`
	Background      string `json:"background"`
	Color_totallose string `json:"color_totallose"`
}
type Model_mlistsipperiodedetail struct {
	Total4d_bayar             int `json:"total4d_bayar"`
	Total3d_bayar             int `json:"total3d_bayar"`
	Total2d_bayar             int `json:"total2d_bayar"`
	Totalcolokbebas_bayar     int `json:"totalcolokbebas_bayar"`
	Totalcolokmacau_bayar     int `json:"totalcolokmacau_bayar"`
	Totalcoloknaga_bayar      int `json:"totalcoloknaga_bayar"`
	Totalcolokjitu_bayar      int `json:"totalcolokjitu_bayar"`
	Total5050umum_bayar       int `json:"total5050umum_bayar"`
	Total5050special_bayar    int `json:"total5050special_bayar"`
	Total5050kombinasi_bayar  int `json:"total5050kombinasi_bayar"`
	Totalmacaukombinasi_bayar int `json:"totalmacaukombinasi_bayar"`
	Totaldasar_bayar          int `json:"totaldasar_bayar"`
	Totalshio_bayar           int `json:"totalshio_bayar"`
	Totalwin_4d               int `json:"totalwin_4d"`
	Totalwin_3d               int `json:"totalwin_3d"`
	Totalwin_2d               int `json:"totalwin_2d"`
	Totalwin_colokbebas       int `json:"totalwin_colokbebas"`
	Totalwin_colokmacau       int `json:"totalwin_colokmacau"`
	Totalwin_coloknaga        int `json:"totalwin_coloknaga"`
	Totalwin_colokjitu        int `json:"totalwin_colokjitu"`
	Totalwin_5050umum         int `json:"totalwin_5050umum"`
	Totalwin_5050special      int `json:"totalwin_5050special"`
	Totalwin_5050kombinasi    int `json:"totalwin_5050kombinasi"`
	Totalwin_macaukombinasi   int `json:"totalwin_macaukombinasi"`
	Totalwin_dasar            int `json:"totalwin_dasar"`
	Totalwin_shio             int `json:"totalwin_shio"`
	Subtotal_bayar            int `json:"subtotal_bayar"`
	Subtotal_winner           int `json:"subtotal_winner"`
	Total_winlose             int `json:"total_winlose"`
}

type Controller_clientToken struct {
	Token    string `json:"token"`
	Hostname string `json:"hostname"`
}
type Controller_clientInit struct {
	Client_Company string `json:"client_company"`
	Hostname       string `json:"hostname"`
}
type Controller_settingcustom struct {
	Status int         `json:"status"`
	Record interface{} `json:"record"`
}
type Controller_clientResult struct {
	Client_Company string `json:"client_company"`
	Pasaran_Code   string `json:"pasaran_code"`
	Hostname       string `json:"hostname"`
}
type Controller_clientResultAll struct {
	Client_Company string `json:"client_company"`
	Hostname       string `json:"hostname"`
}
type Controller_clientConfPasaran struct {
	Client_Company string `json:"client_company"`
	Pasaran_Code   string `json:"pasaran_code"`
	Permainan      string `json:"permainan"`
	Hostname       string `json:"hostname"`
}
type Controller_clientLimitPasaran struct {
	Client_Idinvoice int    `json:"client_idinvoice"`
	Client_Username  string `json:"client_username"`
	Client_Company   string `json:"client_company"`
	Pasaran_Code     string `json:"pasaran_code"`
	Pasaran_Periode  string `json:"pasaran_periode"`
	Permainan        string `json:"permainan"`
	Hostname         string `json:"hostname"`
}
type Controller_clientInvoicePasaran struct {
	Client_Idinvoice int    `json:"client_idinvoice"`
	Client_Username  string `json:"client_username"`
	Client_Company   string `json:"client_company"`
	Pasaran_Code     string `json:"pasaran_code"`
	Pasaran_Periode  string `json:"pasaran_periode"`
	Hostname         string `json:"hostname"`
}
type Controller_clientSlipPeriode struct {
	Client_Username string `json:"client_username"`
	Client_Company  string `json:"client_company"`
	Pasaran_Code    string `json:"pasaran_code"`
	Hostname        string `json:"hostname"`
}
type Controller_clientSlipPeriodeAll struct {
	Client_Username string `json:"client_username"`
	Client_Company  string `json:"client_company"`
	Hostname        string `json:"hostname"`
}
type Controller_clientSlipPeriodeDetail struct {
	Client_Username string `json:"client_username"`
	Client_Company  string `json:"client_company"`
	Idtrxkeluaran   string `json:"idtrxkeluaran"`
	Hostname        string `json:"hostname"`
}
type Controller_clientInvoicePasaranId struct {
	Client_Idinvoice int    `json:"client_idinvoice"`
	Client_Username  string `json:"client_username"`
	Client_Company   string `json:"client_company"`
	Permainan        string `json:"permainan"`
	Hostname         string `json:"hostname"`
}
type Controller_clientSaveTogel struct {
	Client_Username string `json:"client_username"`
	Client_Company  string `json:"client_company"`
	Idtrxkeluaran   string `json:"idtrxkeluaran"`
	Idcomppasaran   string `json:"idcomppasaran"`
	Pasarancode     string `json:"pasarancode"`
	Pasaranperiode  string `json:"pasaranperiode"`
	Devicemember    string `json:"devicemember"`
	Formipaddress   string `json:"formipaddress"`
	Timezone        string `json:"timezone"`
	Totalbayarbet   int    `json:"totalbayarbet"`
	List4d          string `json:"list4d"`
	Hostname        string `json:"hostname"`
}
