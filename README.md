# Flappy Bird Go

Bu proje, klasik Flappy Bird oyununun Go programlama dili kullanılarak yapılmış bir versiyonudur. Oyun, [Ebiten](https://ebiten.org/) 2D oyun motoru kullanılarak geliştirilmiştir.

## Özellikler

- Basit ve eğlenceli oynanış
- Puan sistemi
- Rastgele oluşturulan engeller
- Gerçekçi fizik sistemi
- Oyun sonu ve yeniden başlatma mekanizması

## Gereksinimler

- Go 1.20 veya üzeri
- Ebiten v2 oyun motoru

## Kurulum

1. Go'yu yükleyin (macOS için):
```bash
brew install go
```

2. Projeyi klonlayın:
```bash
git clone [repo-url]
cd flappy-bird-go
```

3. Bağımlılıkları yükleyin:
```bash
go mod tidy
```

## Oyunu Çalıştırma

Oyunu başlatmak için terminal üzerinden şu komutu çalıştırın:
```bash
go run main.go
```

## Nasıl Oynanır?

- SPACE tuşuna basarak kuşu zıplatın
- Yeşil borular arasından geçmeye çalışın
- Her başarılı geçişte 1 puan kazanırsınız
- Borulara veya ekranın üstüne/altına çarparsanız oyun biter
- Oyun bittiğinde SPACE tuşuna basarak yeniden başlatabilirsiniz

## Teknik Detaylar

### Oyun Mekanikleri

1. **Kuş Kontrolü**
   - Yerçekimi etkisi
   - SPACE tuşu ile zıplama
   - Çarpışma kontrolü

2. **Boru Sistemi**
   - Rastgele yüksekliklerde oluşturma
   - Otomatik hareket
   - Ekrandan çıkan boruları temizleme
   - Yeni boruları uygun aralıklarla ekleme

3. **Puan Sistemi**
   - Her boru geçişinde puan artırma
   - Anlık puan gösterimi

### Proje Yapısı

- `main.go`: Tüm oyun kodunu içeren ana dosya
  - `Game` struct'ı: Oyun durumunu yöneten ana yapı
  - `Bird` struct'ı: Kuş karakterinin özelliklerini ve davranışlarını içerir
  - `Pipe` struct'ı: Boru engellerinin özelliklerini ve davranışlarını içerir

## Gelecek Geliştirmeler

- [ ] Sprite'lar ve görseller ekleme
- [ ] Ses efektleri ekleme
- [ ] En yüksek skor sistemi
- [ ] Animasyonlar ekleme
- [ ] Zorluk seviyeleri
- [ ] Farklı karakter seçenekleri

## Gelecek Geliştirmeler

1. Görsel İyileştirmeler
   - Kuş için sprite animasyonları
   - Arka plan görselleri ve paralaks efekti
   - Borular için daha detaylı grafikler
   - Gündüz/gece döngüsü

2. Ses Efektleri
   - Zıplama sesi
   - Puan kazanma sesi
   - Çarpışma sesi
   - Arka plan müziği

3. Oynanış İyileştirmeleri
   - Zorluk seviyeleri (Kolay, Orta, Zor)
   - Güç artırıcılar (yavaş zaman, ekstra can)
   - En yüksek skor tablosu
   - Farklı kuş karakterleri

4. Teknik İyileştirmeler
   - Birim testleri
   - Performans optimizasyonları
   - Mobil platform desteği
   - Ayarlar menüsü (ses, grafik ayarları)

5. Çoklu Platform Desteği
   - Windows build
   - macOS build
   - Linux build
   - Web tarayıcı versiyonu

## Katkıda Bulunma

1. Bu repository'yi fork edin
2. Feature branch'i oluşturun (`git checkout -b feature/amazing-feature`)
3. Değişikliklerinizi commit edin (`git commit -m 'feat: Add some amazing feature'`)
4. Branch'inizi push edin (`git push origin feature/amazing-feature`)
5. Pull Request oluşturun

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır.
