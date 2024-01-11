# webSpider (Go/SQLite)
Go dilinde yazılmış bir web örümcek projesidir. Belirtilen bir web sitesi üzerinde gezinerek tüm yönlendirmeleri (linkleri) tarar ve bu yönlendirmelerin URL, başlık (title) ve içeriğini SQLite veritabanında depolar. Eğer daha önceden taranmış bir URL ile karşılaşılırsa, örümcek bu URL'yi tekrar taramadan diğer yönlendirmeleri araştırmaya devam eder.
### Kullanım:
Proje dizininde terminal veya komut istemcisini açın. "go run webSpider.go" komutunu kullanarak uygulamayı başlatın.Program sizden bir başlangıç URL'si girmenizi isteyecektir. Belirlediğiniz bir web sitesi adresini girin. webSpider, belirtilen site üzerinde dolaşarak linkleri tarayacak ve elde ettiği bilgileri SQLite veritabanına kaydedecektir.
