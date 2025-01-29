# Flappy Bird Go - TODO List

## 1. Asset Sistemi İyileştirmeleri
- [ ] Gerçek kuş sprite'larının eklenmesi
  - [ ] Kuş animasyon kareleri (3 kare)
  - [ ] Kuş rotasyon sistemi
- [ ] Arka plan parallax sistemi
  - [ ] Çoklu arka plan katmanları
  - [ ] Scrolling efekti
- [ ] Boru görselleri
  - [ ] Üst ve alt boru sprite'ları
  - [ ] Boru gölgelendirme efekti

## 2. Ses Sistemi
- [ ] Ses efektleri
  - [ ] Zıplama sesi
  - [ ] Çarpışma sesi
  - [ ] Skor kazanma sesi
- [ ] Arka plan müziği
  - [ ] Oyun müziği
  - [ ] Menü müziği

## 3. Oyun Durumu ve Menü Sistemi
- [ ] Başlangıç menüsü
  - [ ] "Press SPACE to Start" ekranı
  - [ ] Yüksek skor gösterimi
- [ ] Oyun içi UI
  - [ ] Skor gösterimi iyileştirmesi
  - [ ] Zorluk seviyesi göstergesi
- [ ] Game Over ekranı
  - [ ] Son skor
  - [ ] Yüksek skor
  - [ ] Yeniden başlatma seçeneği

## 4. Oyun Mekanikleri
- [ ] Zorluk sistemi
  - [ ] Artan hız
  - [ ] Değişen boru aralıkları
- [ ] Skor sistemi
  - [ ] Yerel yüksek skor kaydı
  - [ ] Başarı puanları
- [ ] Çarpışma sistemi iyileştirmesi
  - [ ] Daha hassas çarpışma kontrolü
  - [ ] Görsel geri bildirim

## 5. Performans İyileştirmeleri
- [ ] Object pooling
  - [ ] Boru havuzu
  - [ ] Efekt havuzu
- [ ] Görüntü önbellekleme
  - [ ] Sprite sheet kullanımı
  - [ ] Render optimizasyonu
- [ ] Bellek yönetimi
  - [ ] Gereksiz tahsislerin azaltılması
  - [ ] GC baskısının düşürülmesi

## 6. Kod Organizasyonu
- [ ] Kod modülerliği
  - [ ] `assets.go` - Asset yönetimi
  - [ ] `score.go` - Skor sistemi
  - [ ] `config.go` - Oyun ayarları
- [ ] Test coverage
  - [ ] Unit testler
  - [ ] Entegrasyon testleri

## 7. Ek Özellikler
- [ ] Farklı kuş karakterleri
- [ ] Gündüz/gece döngüsü
- [ ] Hava efektleri (yağmur, kar)
- [ ] Achievements sistemi

## Öncelik Sırası
1. Temel oyun mekaniklerinin iyileştirilmesi
2. Asset sisteminin tamamlanması
3. Ses sisteminin eklenmesi
4. UI ve menü sisteminin geliştirilmesi
5. Performans optimizasyonları
6. Ek özelliklerin eklenmesi
