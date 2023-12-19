
# GO & MYSQL API ÖRNEĞİ

Bu projenin ne yaptığı ve kimin için olduğu hakkında kısa bir açıklama



## API Kullanımı
Bu API Golang ve MYSQL kullanılarak oluşturulmuştur. Temel amacı string bir title değeri olan database bağlantılı bir API kaynağı oluşturulmaktır. 

                | (API testlerinde Postman kullanıldı) |

Not: İstenilirse model klasörü altında bulunan struct yapısı kullanılarak projenin yapısı değitirilebilir.

Boş zamanımda main.go da bulunan dosyaları parçalara ayırmak istiyorum. 
#### Tüm öğeleri getir

```http
  GET localhost:5555/posts
```

#### Id değerine göre

```http
  GET localhost:5555/posts/3
```
#### Yeni veri ekleme

```http
  POST localhost:5555/posts/3
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `title`      | `string` | **Gerekli**. Eklenecek yeni veri |


#### Veri güncelleme

```http
  PUT localhost:5555/posts/3
```
| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `title`      | `string` | **Gerekli**. Güncellenecek yeni veri |


#### Veri silme

```http
  DELETE localhost:5555/posts/3
```
