# Cyber & Pastel Design System Style Guide

A comprehensive design system featuring two distinctive palettes with modern typography for digital interfaces.

## Table of Contents
- [Color Palettes](#color-palettes)
- [Typography System](#typography-system)
- [Implementation Guide](#implementation-guide)
- [Usage Examples](#usage-examples)
- [Best Practices](#best-practices)

---

## Color Palettes

### 🌐 Cyber Black Palette
**Perfect for:** Gaming interfaces, tech dashboards, cyberpunk themes, high-energy digital experiences

#### Base Colors
| Color Name | Hex Code | Usage |
|------------|----------|--------|
| **Void Black** | `#000000` | Primary background, deep shadows, maximum contrast base |
| **Carbon Black** | `#0a0a0a` | Secondary background, subtle elevation |
| **Cyber Gray** | `#1a1a1a` | Card backgrounds, modal overlays |
| **Pure White** | `#ffffff` | Primary text, high contrast elements |

#### Accent Colors
| Color Name | Hex Code | Usage |
|------------|----------|--------|
| **Matrix Green** | `#00ff00` | Success states, online status, completion |
| **Neon Pink** | `#ff0080` | Error states, critical alerts, danger |
| **Electric Cyan** | `#00ffff` | Processing states, loading, active elements |
| **Laser Yellow** | `#ffff00` | Warning states, caution, pending actions |

#### CSS Variables
```css
:root {
  /* Cyber Black Palette */
  --cyber-void-black: #000000;
  --cyber-carbon-black: #0a0a0a;
  --cyber-gray: #1a1a1a;
  --cyber-white: #ffffff;
  --cyber-matrix-green: #00ff00;
  --cyber-neon-pink: #ff0080;
  --cyber-electric-cyan: #00ffff;
  --cyber-laser-yellow: #ffff00;
}
```

### 🌸 Pastel Black Palette
**Perfect for:** Mobile apps, productivity tools, wellness apps, modern websites, professional interfaces

#### Base Colors
| Color Name | Hex Code | Usage |
|------------|----------|--------|
| **Midnight Black** | `#000000` | Primary background, deep contrast foundation |
| **Charcoal** | `#0f0f0f` | Secondary surfaces, subtle depth |
| **Soft Gray** | `#1e1e1e` | Card backgrounds, gentle elevation |
| **Soft White** | `#f8f8f8` | Primary text, comfortable reading |

#### Accent Colors
| Color Name | Hex Code | Usage |
|------------|----------|--------|
| **Mint Green** | `#a8e6a3` | Success states, positive feedback, completion |
| **Blush Pink** | `#ffb3ba` | Error states, gentle warnings, attention |
| **Sky Blue** | `#b3e5fc` | Processing states, information, calm activity |
| **Cream Yellow** | `#fff9c4` | Warning states, caution, pending review |

#### CSS Variables
```css
:root {
  /* Pastel Black Palette */
  --pastel-midnight-black: #000000;
  --pastel-charcoal: #0f0f0f;
  --pastel-soft-gray: #1e1e1e;
  --pastel-soft-white: #f8f8f8;
  --pastel-mint-green: #a8e6a3;
  --pastel-blush-pink: #ffb3ba;
  --pastel-sky-blue: #b3e5fc;
  --pastel-cream-yellow: #fff9c4;
}
```

---

## Typography System

### Font Selection
- **UI Font**: Inter Nerd Font - Clean, modern, highly readable
- **Monospace Font**: Space Mono Nerd Font - Unique personality, perfect for code

### Font Stack Definition
```css
:root {
  /* Font Stacks */
  --font-ui: 'Inter', 'Inter Nerd Font', 'Segoe UI', 'Roboto', sans-serif;
  --font-mono: 'Space Mono', 'Space Mono Nerd Font', 'JetBrains Mono', monospace;
  
  /* Font Weights */
  --font-light: 300;
  --font-regular: 400;
  --font-medium: 500;
  --font-semibold: 600;
  --font-bold: 700;
  
  /* Font Sizes */
  --text-xs: 0.75rem;    /* 12px */
  --text-sm: 0.875rem;   /* 14px */
  --text-base: 1rem;     /* 16px */
  --text-lg: 1.125rem;   /* 18px */
  --text-xl: 1.25rem;    /* 20px */
  --text-2xl: 1.5rem;    /* 24px */
  --text-3xl: 1.875rem;  /* 30px */
  --text-4xl: 2.25rem;   /* 36px */
}
```

### Font Loading
#### CDN Method (Fallback)
```css
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Space+Mono:wght@400;700&display=swap');
```

#### Self-Hosted Method (Recommended)
```css
@font-face {
  font-family: 'Inter Nerd Font';
  src: url('./fonts/InterNerdFont-Regular.woff2') format('woff2');
  font-weight: 400;
  font-style: normal;
}

@font-face {
  font-family: 'Space Mono Nerd Font';
  src: url('./fonts/SpaceMonoNerdFont-Regular.woff2') format('woff2');
  font-weight: 400;
  font-style: normal;
}
```

### Typography Scale
| Size | Rem | Pixels | Usage |
|------|-----|---------|--------|
| **XS** | 0.75rem | 12px | Tiny details, footnotes |
| **SM** | 0.875rem | 14px | Small labels, metadata |
| **Base** | 1rem | 16px | Body text, regular content |
| **LG** | 1.125rem | 18px | Subheadings, emphasis |
| **XL** | 1.25rem | 20px | Card titles, section headers |
| **2XL** | 1.5rem | 24px | Page titles, main headings |
| **3XL** | 1.875rem | 30px | Display headings |
| **4XL** | 2.25rem | 36px | Hero text, major displays |

---

## Implementation Guide

### Semantic CSS Classes
```css
/* Font Families */
.font-ui { font-family: var(--font-ui); }
.font-mono { font-family: var(--font-mono); }

/* Font Weights */
.weight-light { font-weight: var(--font-light); }
.weight-regular { font-weight: var(--font-regular); }
.weight-medium { font-weight: var(--font-medium); }
.weight-semibold { font-weight: var(--font-semibold); }
.weight-bold { font-weight: var(--font-bold); }

/* Font Sizes */
.text-xs { font-size: var(--text-xs); }
.text-sm { font-size: var(--text-sm); }
.text-base { font-size: var(--text-base); }
.text-lg { font-size: var(--text-lg); }
.text-xl { font-size: var(--text-xl); }
.text-2xl { font-size: var(--text-2xl); }
.text-3xl { font-size: var(--text-3xl); }
.text-4xl { font-size: var(--text-4xl); }

/* Component Classes */
.heading { 
  font-family: var(--font-ui); 
  font-weight: var(--font-semibold); 
}

.body-text { 
  font-family: var(--font-ui); 
  font-weight: var(--font-regular); 
}

.code-block { 
  font-family: var(--font-mono); 
  font-weight: var(--font-regular); 
}

.terminal { 
  font-family: var(--font-mono); 
  background: var(--cyber-carbon-black); 
  color: var(--cyber-white); 
}
```

### Color Utility Classes
```css
/* Cyber Theme */
.cyber-bg-primary { background-color: var(--cyber-void-black); }
.cyber-bg-secondary { background-color: var(--cyber-carbon-black); }
.cyber-bg-surface { background-color: var(--cyber-gray); }
.cyber-text-primary { color: var(--cyber-white); }
.cyber-text-success { color: var(--cyber-matrix-green); }
.cyber-text-error { color: var(--cyber-neon-pink); }
.cyber-text-warning { color: var(--cyber-laser-yellow); }
.cyber-text-info { color: var(--cyber-electric-cyan); }

/* Pastel Theme */
.pastel-bg-primary { background-color: var(--pastel-midnight-black); }
.pastel-bg-secondary { background-color: var(--pastel-charcoal); }
.pastel-bg-surface { background-color: var(--pastel-soft-gray); }
.pastel-text-primary { color: var(--pastel-soft-white); }
.pastel-text-success { color: var(--pastel-mint-green); }
.pastel-text-error { color: var(--pastel-blush-pink); }
.pastel-text-warning { color: var(--pastel-cream-yellow); }
.pastel-text-info { color: var(--pastel-sky-blue); }
```

---

## Usage Examples

### Status Indicators

#### Cyber Theme
```html
<!-- Online Status -->
<div class="status-indicator">
  <span class="status-dot" style="background: var(--cyber-matrix-green);"></span>
  <span class="font-ui text-sm cyber-text-primary">System Online</span>
</div>

<!-- Error Status -->
<div class="status-indicator">
  <span class="status-dot" style="background: var(--cyber-neon-pink);"></span>
  <span class="font-ui text-sm cyber-text-primary">Connection Failed</span>
</div>

<!-- Processing Status -->
<div class="status-indicator">
  <span class="status-dot" style="background: var(--cyber-electric-cyan);"></span>
  <span class="font-ui text-sm cyber-text-primary">Processing...</span>
</div>
```

#### Pastel Theme
```html
<!-- Success Status -->
<div class="status-indicator">
  <span class="status-dot" style="background: var(--pastel-mint-green);"></span>
  <span class="font-ui text-sm pastel-text-primary">Task Complete</span>
</div>

<!-- Warning Status -->
<div class="status-indicator">
  <span class="status-dot" style="background: var(--pastel-cream-yellow);"></span>
  <span class="font-ui text-sm pastel-text-primary">Needs Review</span>
</div>
```

### Button Components

#### Cyber Theme Buttons
```html
<!-- Success Button -->
<button class="btn cyber-success">
  <span class="font-ui text-sm weight-medium">Execute Command</span>
</button>

<!-- Error Button -->
<button class="btn cyber-error">
  <span class="font-ui text-sm weight-medium">Terminate Process</span>
</button>

<!-- Info Button -->
<button class="btn cyber-info">
  <span class="font-ui text-sm weight-medium">Scan System</span>
</button>
```

```css
.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 0.375rem;
  cursor: pointer;
  transition: all 0.2s ease;
}

.cyber-success {
  background: var(--cyber-matrix-green);
  color: var(--cyber-void-black);
}

.cyber-error {
  background: var(--cyber-neon-pink);
  color: var(--cyber-white);
}

.cyber-info {
  background: var(--cyber-electric-cyan);
  color: var(--cyber-void-black);
}
```

### Terminal/Code Blocks

#### Cyber Terminal
```html
<div class="terminal-window">
  <div class="terminal-header">
    <span class="font-mono text-sm weight-bold cyber-text-info">SYSTEM TERMINAL</span>
  </div>
  <div class="terminal-content">
    <div class="font-mono text-sm cyber-text-info">$ system.status --verbose</div>
    <div class="font-mono text-sm cyber-text-success">✓ Connection established</div>
    <div class="font-mono text-sm cyber-text-warning">⚠ Memory usage: 84%</div>
    <div class="font-mono text-sm cyber-text-error">✗ Critical error detected</div>
  </div>
</div>
```

### Card Components

#### Pastel Card
```html
<div class="card pastel-theme">
  <div class="card-header">
    <h3 class="font-ui text-xl weight-semibold pastel-text-primary">
      Project Status
    </h3>
  </div>
  <div class="card-body">
    <p class="font-ui text-base weight-regular pastel-text-primary">
      Your project is running smoothly with no issues detected.
    </p>
    <div class="status-grid">
      <div class="status-item">
        <span class="font-ui text-sm weight-medium pastel-text-success">
          ✓ All systems operational
        </span>
      </div>
    </div>
  </div>
</div>
```

---

## Best Practices

### Color Usage Guidelines

#### Cyber Theme
- **Use high contrast** for maximum visibility
- **Neon accents sparingly** - they should pop, not overwhelm
- **Matrix Green** for positive actions and success states
- **Neon Pink** for critical errors and destructive actions
- **Electric Cyan** for interactive elements and processing states
- **Laser Yellow** for warnings and caution states

#### Pastel Theme
- **Maintain readability** with sufficient contrast
- **Soft accents** should feel gentle and non-intrusive
- **Mint Green** for positive feedback and completion
- **Blush Pink** for gentle errors and attention
- **Sky Blue** for information and calm interactions
- **Cream Yellow** for subtle warnings

### Typography Guidelines

#### Font Selection
- **Inter Nerd Font** for all UI elements, buttons, headings, and body text
- **Space Mono Nerd Font** for code, terminal output, data display, and technical content
- **Never mix** more than two font families in a single interface

#### Hierarchy
1. **Display Text** (3XL-4XL, Bold) - Hero sections, major headings
2. **Headings** (XL-2XL, Semibold) - Section titles, card headers
3. **Body Text** (Base-LG, Regular) - Main content, descriptions
4. **Small Text** (SM-XS, Medium) - Labels, metadata, captions

#### Line Height
- **Display Text**: 1.1-1.2
- **Headings**: 1.2-1.3
- **Body Text**: 1.4-1.6
- **Code/Terminal**: 1.4

### Accessibility Considerations

#### Color Contrast
- **Cyber Theme**: High contrast ratios (7:1 or higher)
- **Pastel Theme**: Maintain WCAG AA compliance (4.5:1 minimum)
- **Never rely on color alone** for important information

#### Font Accessibility
- **Minimum 16px** for body text
- **Maximum line length** of 75 characters
- **Sufficient spacing** between interactive elements (44px minimum)

### Implementation Tips

#### Performance
- **Load fonts efficiently** using font-display: swap
- **Preload critical fonts** for faster rendering
- **Use system fonts** as fallbacks

#### Maintenance
- **Use CSS custom properties** for easy theme switching
- **Document color meanings** and usage contexts
- **Test both themes** in different lighting conditions

#### Responsiveness
- **Scale typography** appropriately for different screen sizes
- **Adjust contrast** for different viewing environments
- **Test on various devices** and browsers

---

## File Structure

```
styles/
├── tokens/
│   ├── colors.css          # Color variables
│   ├── typography.css      # Font and text variables
│   └── spacing.css         # Spacing and layout variables
├── components/
│   ├── buttons.css         # Button styles
│   ├── cards.css           # Card components
│   ├── forms.css           # Form elements
│   └── status.css          # Status indicators
├── themes/
│   ├── cyber.css           # Cyber theme overrides
│   └── pastel.css          # Pastel theme overrides
└── main.css               # Main stylesheet
```

---

## Resources

### Font Downloads
- **Inter Nerd Font**: [Nerd Fonts Repository](https://github.com/ryanoasis/nerd-fonts)
- **Space Mono Nerd Font**: [Nerd Fonts Repository](https://github.com/ryanoasis/nerd-fonts)

### Color Tools
- **Contrast Checker**: [WebAIM Contrast Checker](https://webaim.org/resources/contrastchecker/)
- **Accessibility**: [WCAG Guidelines](https://www.w3.org/WAI/WCAG21/quickref/)

### Browser Support
- **CSS Custom Properties**: IE11+ (with polyfill)
- **Font Loading**: All modern browsers
- **Fallback Fonts**: Ensure graceful degradation

---

*This style guide provides a comprehensive foundation for building consistent, accessible, and visually appealing interfaces using the Cyber and Pastel design systems.*