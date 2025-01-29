# Flappy Bird Go

A Flappy Bird clone written in Go using Ebitengine.

## Project Structure

```
flappy-bird-go/
├── cmd/
│   └── flappybird/         # Main application entry point
├── internal/
│   ├── game/               # Game logic
│   ├── assets/
│   │   ├── sprites/        # Game sprites and images
│   │   └── audio/          # Sound effects and music
│   └── config/             # Game configuration
├── assets/                 # Raw assets
└── tools/                  # Development tools
```

## Development Roadmap

### Current Sprint: Visual Improvements
- [ ] Implement sprite animations for bird
- [ ] Add background graphics with parallax effect
- [ ] Enhance pipe graphics
- [ ] Add day/night cycle

### Upcoming Features

1. Sound System
   - [ ] Jump sound effect
   - [ ] Score sound effect
   - [ ] Collision sound effect
   - [ ] Background music

2. Gameplay Enhancements
   - [ ] Difficulty levels (Easy, Medium, Hard)
   - [ ] Power-ups (slow time, extra life)
   - [ ] High score system
   - [ ] Multiple bird characters

3. Technical Improvements
   - [ ] Unit tests
   - [ ] Performance optimizations
   - [ ] Mobile platform support
   - [ ] Settings menu (sound, graphics)

4. Multi-Platform Support
   - [ ] Windows build
   - [ ] macOS build
   - [ ] Linux build
   - [ ] Web browser version

## Getting Started

### Prerequisites
- Go 1.20 or higher
- Ebitengine dependencies

### Installation
```bash
git clone https://github.com/yourusername/flappy-bird-go
cd flappy-bird-go
go mod download
```

### Running the Game
```bash
go run main.go
```

## Controls
- Press SPACE to make the bird jump
- Press SPACE to restart when game over

## Özellikler

- Basit ve eğlenceli oynanış
- Puan sistemi
- Rastgele oluşturulan engeller
- Gerçekçi fizik sistemi
- Oyun sonu ve yeniden başlatma mekanizması

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

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır.
