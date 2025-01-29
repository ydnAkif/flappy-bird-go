# Flappy Bird Go - Proje Yapısı

```
flappy-bird-go/
├── assets/                 # Oyun varlıkları
│   ├── images/            # Görsel varlıklar
│   │   ├── bird/         # Kuş animasyon kareleri
│   │   ├── background/    # Arka plan katmanları
│   │   └── pipes/        # Boru görselleri
│   ├── audio/            # Ses varlıkları
│   │   ├── effects/      # Ses efektleri
│   │   └── music/        # Arka plan müzikleri
│   └── fonts/            # Oyun fontları
│
├── cmd/                   # Uygulama giriş noktası
│   └── flappybird/
│       └── main.go       # Ana uygulama dosyası
│
├── internal/             # İç paketler
│   ├── game/            # Oyun mantığı
│   │   ├── game.go      # Ana oyun yapısı
│   │   ├── bird.go      # Kuş mantığı
│   │   ├── pipe.go      # Boru mantığı
│   │   └── collision.go # Çarpışma sistemi
│   │
│   ├── assets/          # Asset yönetimi
│   │   ├── loader.go    # Asset yükleme sistemi
│   │   └── cache.go     # Asset önbellekleme
│   │
│   ├── audio/          # Ses sistemi
│   │   ├── manager.go  # Ses yönetimi
│   │   └── player.go   # Ses çalma
│   │
│   └── ui/             # Kullanıcı arayüzü
│       ├── menu.go     # Menü sistemi
│       └── hud.go      # Oyun içi arayüz
│
├── pkg/                 # Dışa açık paketler
│   ├── animation/      # Animasyon sistemi
│   └── physics/        # Fizik hesaplamaları
│
├── configs/            # Yapılandırma dosyaları
│   └── game.json      # Oyun ayarları
│
└── tests/             # Test dosyaları
    ├── game_test.go
    └── collision_test.go

# Kök Dizindeki Dosyalar
├── .gitignore         # Git yoksayma listesi
├── go.mod            # Go modül tanımı
├── go.sum            # Go bağımlılık hash'leri
├── LICENSE           # Lisans dosyası
├── README.md         # Proje dokümantasyonu
└── TODO.md           # Yapılacaklar listesi
```

## Paket Sorumlulukları

### cmd/
- Uygulama başlangıç noktası
- Temel yapılandırma
- Hata yönetimi

### internal/game/
- Oyun durumu yönetimi
- Fizik hesaplamaları
- Çarpışma tespiti
- Skor sistemi

### internal/assets/
- Görsel ve ses dosyalarının yüklenmesi
- Asset önbellekleme
- Kaynak yönetimi

### internal/audio/
- Ses efektleri yönetimi
- Müzik çalma
- Ses kontrolleri

### internal/ui/
- Menü sistemi
- Skor gösterimi
- Oyun içi arayüz

### pkg/
- Yeniden kullanılabilir bileşenler
- Genel amaçlı yardımcı fonksiyonlar

## Geliştirme Kuralları

1. Her paket kendi sorumluluğuna odaklanmalı
2. Paketler arası bağımlılıklar minimize edilmeli
3. Testler ilgili kodla aynı pakette olmalı
4. Dış dünyaya açık API'lar pkg/ altında olmalı
5. Asset dosyaları git LFS kullanılarak yönetilmeli
