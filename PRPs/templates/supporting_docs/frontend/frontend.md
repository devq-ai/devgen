# .rules_ui - Complete Design System

A unified design system for modern applications with cohesive color palettes, typography, and component guidelines.

## Table of Contents
- [Design Philosophy](#design-philosophy)
- [Color System](#color-system)
- [Typography System](#typography-system)
- [Application-Specific Guidelines](#application-specific-guidelines)
- [Component Library](#component-library)
- [Animation System](#animation-system)
- [Implementation Guide](#implementation-guide)

---

## Design Philosophy

### Core Principles
- **Consistency**: Unified visual language across all platforms
- **Accessibility**: WCAG AA compliance minimum, AAA preferred
- **Performance**: Optimized for fast loading and smooth interactions
- **Scalability**: Design tokens that scale from mobile to desktop
- **Developer Experience**: Clear guidelines and reusable components

### Theme Strategy
- **Cyber Black**: High-energy, technical interfaces (NextJS frontend, CLI)
- **Pastel Black**: Calm, professional dashboards (Python panel)
- **Existing Palettes**: Maintained for legacy compatibility

---

## Color System

### Primary Color Palettes

#### 🌐 Cyber Black Palette
**Recommended for**: NextJS Frontend, CLI Applications, Developer Tools

```css
:root {
  /* Base Colors */
  --cyber-void-black: #000000;
  --cyber-carbon-black: #0a0a0a;
  --cyber-cyber-gray: #1a1a1a;
  --cyber-pure-white: #ffffff;
  
  /* Accent Colors */
  --cyber-matrix-green: #00ff00;
  --cyber-neon-pink: #ff0080;
  --cyber-electric-cyan: #00ffff;
  --cyber-laser-yellow: #ffff00;
  
  /* Semantic Colors */
  --cyber-success: var(--cyber-matrix-green);
  --cyber-error: var(--cyber-neon-pink);
  --cyber-warning: var(--cyber-laser-yellow);
  --cyber-info: var(--cyber-electric-cyan);
  --cyber-primary: var(--cyber-electric-cyan);
  --cyber-secondary: var(--cyber-cyber-gray);
}
```

#### 🌸 Pastel Black Palette
**Recommended for**: Python Panel Dashboard, Data Visualization, Admin Interfaces

```css
:root {
  /* Base Colors */
  --pastel-midnight-black: #000000;
  --pastel-charcoal: #0f0f0f;
  --pastel-soft-gray: #1e1e1e;
  --pastel-soft-white: #f8f8f8;
  
  /* Accent Colors */
  --pastel-mint-green: #a8e6a3;
  --pastel-blush-pink: #ffb3ba;
  --pastel-sky-blue: #b3e5fc;
  --pastel-cream-yellow: #fff9c4;
  --pastel-lavender: #d4c5f9;
  --pastel-peach: #ffccb3;
  
  /* Semantic Colors */
  --pastel-success: var(--pastel-mint-green);
  --pastel-error: var(--pastel-blush-pink);
  --pastel-warning: var(--pastel-cream-yellow);
  --pastel-info: var(--pastel-sky-blue);
  --pastel-primary: var(--pastel-sky-blue);
  --pastel-secondary: var(--pastel-lavender);
}
```

#### 🎨 Extended Palette Integration
**Cohesive with existing color schemes**

```css
:root {
  /* Neutral Scale */
  --neutral-50: #fafafa;
  --neutral-100: #f5f5f5;
  --neutral-200: #e5e5e5;
  --neutral-300: #d4d4d4;
  --neutral-400: #a3a3a3;
  --neutral-500: #737373;
  --neutral-600: #525252;
  --neutral-700: #404040;
  --neutral-800: #262626;
  --neutral-900: #171717;
  --neutral-950: #0a0a0a;
  
  /* Brand Colors */
  --brand-primary: #3b82f6;
  --brand-secondary: #8b5cf6;
  --brand-accent: #06b6d4;
  
  /* Surface Colors */
  --surface-primary: var(--neutral-900);
  --surface-secondary: var(--neutral-800);
  --surface-tertiary: var(--neutral-700);
  --surface-overlay: rgba(0, 0, 0, 0.8);
  
  /* Text Colors */
  --text-primary: var(--neutral-50);
  --text-secondary: var(--neutral-300);
  --text-tertiary: var(--neutral-400);
  --text-inverse: var(--neutral-900);
}
```

---

## Typography System

### Font Selection Strategy

#### Primary Fonts
- **UI Font**: Inter Nerd Font - Modern, readable, excellent for interfaces
- **Code Font**: Space Mono Nerd Font - Distinctive, perfect for terminals and code
- **Display Font**: Inter - Clean hierarchy for headings

#### Font Loading & Fallbacks
```css
:root {
  /* Font Stacks */
  --font-ui: 'Inter', 'Inter Nerd Font', 'Segoe UI', system-ui, sans-serif;
  --font-mono: 'Space Mono', 'Space Mono Nerd Font', 'JetBrains Mono', 'Fira Code', monospace;
  --font-display: 'Inter', 'Inter Nerd Font', 'Segoe UI', system-ui, sans-serif;
  
  /* Font Weights */
  --font-thin: 100;
  --font-light: 300;
  --font-regular: 400;
  --font-medium: 500;
  --font-semibold: 600;
  --font-bold: 700;
  --font-extrabold: 800;
  --font-black: 900;
  
  /* Font Sizes */
  --text-xs: 0.75rem;      /* 12px */
  --text-sm: 0.875rem;     /* 14px */
  --text-base: 1rem;       /* 16px */
  --text-lg: 1.125rem;     /* 18px */
  --text-xl: 1.25rem;      /* 20px */
  --text-2xl: 1.5rem;      /* 24px */
  --text-3xl: 1.875rem;    /* 30px */
  --text-4xl: 2.25rem;     /* 36px */
  --text-5xl: 3rem;        /* 48px */
  --text-6xl: 3.75rem;     /* 60px */
  
  /* Line Heights */
  --leading-tight: 1.25;
  --leading-normal: 1.5;
  --leading-relaxed: 1.75;
  
  /* Letter Spacing */
  --tracking-tight: -0.025em;
  --tracking-normal: 0;
  --tracking-wide: 0.05em;
}
```

### Typography Scale & Hierarchy
```css
/* Semantic Typography Classes */
.display-xl {
  font-size: var(--text-6xl);
  font-weight: var(--font-bold);
  line-height: var(--leading-tight);
  letter-spacing: var(--tracking-tight);
}

.display-lg {
  font-size: var(--text-5xl);
  font-weight: var(--font-bold);
  line-height: var(--leading-tight);
}

.heading-xl {
  font-size: var(--text-4xl);
  font-weight: var(--font-semibold);
  line-height: var(--leading-tight);
}

.heading-lg {
  font-size: var(--text-3xl);
  font-weight: var(--font-semibold);
  line-height: var(--leading-normal);
}

.heading-md {
  font-size: var(--text-2xl);
  font-weight: var(--font-semibold);
  line-height: var(--leading-normal);
}

.heading-sm {
  font-size: var(--text-xl);
  font-weight: var(--font-medium);
  line-height: var(--leading-normal);
}

.body-lg {
  font-size: var(--text-lg);
  font-weight: var(--font-regular);
  line-height: var(--leading-relaxed);
}

.body-md {
  font-size: var(--text-base);
  font-weight: var(--font-regular);
  line-height: var(--leading-normal);
}

.body-sm {
  font-size: var(--text-sm);
  font-weight: var(--font-regular);
  line-height: var(--leading-normal);
}

.caption {
  font-size: var(--text-xs);
  font-weight: var(--font-medium);
  line-height: var(--leading-normal);
  text-transform: uppercase;
  letter-spacing: var(--tracking-wide);
}
```

---

## Application-Specific Guidelines

### 🚀 NextJS Frontend
**Theme**: Cyber Black Palette
**Focus**: High-energy, modern web application

#### Color Usage
```css
/* NextJS Theme Variables */
:root {
  --nextjs-bg-primary: var(--cyber-void-black);
  --nextjs-bg-secondary: var(--cyber-carbon-black);
  --nextjs-bg-surface: var(--cyber-cyber-gray);
  --nextjs-text-primary: var(--cyber-pure-white);
  --nextjs-accent-primary: var(--cyber-electric-cyan);
  --nextjs-accent-secondary: var(--cyber-matrix-green);
  --nextjs-success: var(--cyber-matrix-green);
  --nextjs-error: var(--cyber-neon-pink);
  --nextjs-warning: var(--cyber-laser-yellow);
  --nextjs-info: var(--cyber-electric-cyan);
}
```

#### Typography
- **Primary**: Inter Nerd Font for all UI elements
- **Accents**: Space Mono Nerd Font for code snippets, data display
- **Hierarchy**: Bold headings, medium body text, regular captions

#### Recommended shadcn/ui Components
```typescript
// Essential Components for NextJS
import {
  Button,
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
  Badge,
  Input,
  Label,
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  Popover,
  PopoverContent,
  PopoverTrigger,
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
  Switch,
  Textarea,
  Toast,
  Toaster,
  Progress,
  Skeleton,
  Avatar,
  AvatarFallback,
  AvatarImage,
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
} from "@/components/ui"
```

### 📊 Python Panel Dashboard
**Theme**: Pastel Black Palette
**Focus**: Data visualization, analytics, professional interface

#### Color Usage
```css
/* Python Dashboard Theme Variables */
:root {
  --dashboard-bg-primary: var(--pastel-midnight-black);
  --dashboard-bg-secondary: var(--pastel-charcoal);
  --dashboard-bg-surface: var(--pastel-soft-gray);
  --dashboard-text-primary: var(--pastel-soft-white);
  --dashboard-accent-primary: var(--pastel-sky-blue);
  --dashboard-accent-secondary: var(--pastel-lavender);
  --dashboard-success: var(--pastel-mint-green);
  --dashboard-error: var(--pastel-blush-pink);
  --dashboard-warning: var(--pastel-cream-yellow);
  --dashboard-info: var(--pastel-sky-blue);
  --dashboard-chart-1: var(--pastel-sky-blue);
  --dashboard-chart-2: var(--pastel-mint-green);
  --dashboard-chart-3: var(--pastel-lavender);
  --dashboard-chart-4: var(--pastel-peach);
  --dashboard-chart-5: var(--pastel-cream-yellow);
  --dashboard-chart-6: var(--pastel-blush-pink);
}
```

#### Typography
- **Primary**: Inter Nerd Font for readability
- **Data**: Space Mono Nerd Font for metrics and numbers
- **Hierarchy**: Medium headings, regular body text, light captions

#### Panel-Specific Styling
```python
# Panel CSS Integration
pn.config.raw_css = """
:root {
  --panel-primary: var(--dashboard-bg-primary);
  --panel-secondary: var(--dashboard-bg-secondary);
  --panel-surface: var(--dashboard-bg-surface);
  --panel-text: var(--dashboard-text-primary);
  --panel-accent: var(--dashboard-accent-primary);
}

.bk-root {
  font-family: var(--font-ui);
  background: var(--panel-primary);
  color: var(--panel-text);
}

.panel-widget-box {
  background: var(--panel-surface);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  padding: 1rem;
}
"""
```

### 💻 CLI Applications
**Theme**: Cyber Black Palette
**Focus**: Terminal interface, developer tools

#### Color Usage
```css
/* CLI Theme Variables */
:root {
  --cli-bg-primary: var(--cyber-void-black);
  --cli-bg-secondary: var(--cyber-carbon-black);
  --cli-text-primary: var(--cyber-pure-white);
  --cli-text-secondary: var(--neutral-300);
  --cli-success: var(--cyber-matrix-green);
  --cli-error: var(--cyber-neon-pink);
  --cli-warning: var(--cyber-laser-yellow);
  --cli-info: var(--cyber-electric-cyan);
  --cli-prompt: var(--cyber-electric-cyan);
  --cli-command: var(--cyber-pure-white);
  --cli-output: var(--neutral-300);
  --cli-accent: var(--cyber-matrix-green);
}
```

#### Typography
- **Primary**: Space Mono Nerd Font for all terminal text
- **Hierarchy**: Regular weight throughout, color for emphasis
- **Formatting**: Consistent spacing, clear readability

#### CLI Color Codes
```python
# ANSI Color Codes for CLI
CLI_COLORS = {
    'success': '\033[92m',     # Matrix Green
    'error': '\033[91m',       # Neon Pink  
    'warning': '\033[93m',     # Laser Yellow
    'info': '\033[96m',        # Electric Cyan
    'prompt': '\033[96m',      # Electric Cyan
    'reset': '\033[0m',        # Reset
    'bold': '\033[1m',         # Bold
    'dim': '\033[2m',          # Dim
}
```

---

## Component Library

### shadcn/ui Component Selection & Customization

#### Core Components
```typescript
// components/ui/button.tsx
const buttonVariants = cva(
  "inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors focus-visible:outline-none disabled:pointer-events-none disabled:opacity-50",
  {
    variants: {
      variant: {
        // Cyber Theme Variants
        'cyber-primary': 'bg-cyber-electric-cyan text-cyber-void-black hover:bg-cyber-electric-cyan/80',
        'cyber-success': 'bg-cyber-matrix-green text-cyber-void-black hover:bg-cyber-matrix-green/80',
        'cyber-error': 'bg-cyber-neon-pink text-cyber-pure-white hover:bg-cyber-neon-pink/80',
        'cyber-warning': 'bg-cyber-laser-yellow text-cyber-void-black hover:bg-cyber-laser-yellow/80',
        'cyber-outline': 'border border-cyber-electric-cyan text-cyber-electric-cyan hover:bg-cyber-electric-cyan hover:text-cyber-void-black',
        
        // Pastel Theme Variants
        'pastel-primary': 'bg-pastel-sky-blue text-pastel-midnight-black hover:bg-pastel-sky-blue/80',
        'pastel-success': 'bg-pastel-mint-green text-pastel-midnight-black hover:bg-pastel-mint-green/80',
        'pastel-error': 'bg-pastel-blush-pink text-pastel-midnight-black hover:bg-pastel-blush-pink/80',
        'pastel-warning': 'bg-pastel-cream-yellow text-pastel-midnight-black hover:bg-pastel-cream-yellow/80',
        'pastel-outline': 'border border-pastel-sky-blue text-pastel-sky-blue hover:bg-pastel-sky-blue hover:text-pastel-midnight-black',
      },
      size: {
        default: "h-10 px-4 py-2",
        sm: "h-9 rounded-md px-3",
        lg: "h-11 rounded-md px-8",
        icon: "h-10 w-10",
      },
    },
    defaultVariants: {
      variant: "cyber-primary",
      size: "default",
    },
  }
)
```

#### Data Display Components
```typescript
// For Dashboard/Analytics
const dashboardComponents = [
  'Card',
  'Table',
  'Badge',
  'Progress',
  'Skeleton',
  'Chart', // Custom chart wrapper
  'Metric', // Custom metric display
  'KPI',    // Custom KPI component
]

// For NextJS Frontend
const frontendComponents = [
  'NavigationMenu',
  'Sheet',
  'Dialog',
  'Command',
  'Combobox',
  'DataTable',
  'Form',
  'Input',
  'Select',
  'Tabs',
  'Toast',
]
```

#### Custom Component Extensions
```typescript
// components/ui/terminal.tsx
interface TerminalProps {
  commands: Array<{
    input: string
    output: string
    type: 'success' | 'error' | 'warning' | 'info'
  }>
}

export function Terminal({ commands }: TerminalProps) {
  return (
    <div className="bg-cyber-void-black border border-cyber-cyber-gray rounded-lg p-4 font-mono text-sm">
      {commands.map((cmd, i) => (
        <div key={i} className="mb-2">
          <div className="text-cyber-electric-cyan">
            <span className="text-cyber-matrix-green">$</span> {cmd.input}
          </div>
          <div className={`text-cyber-${cmd.type} ml-2`}>
            {cmd.output}
          </div>
        </div>
      ))}
    </div>
  )
}
```

```typescript
// components/ui/metric-card.tsx
interface MetricCardProps {
  title: string
  value: string | number
  change?: {
    value: number
    type: 'increase' | 'decrease'
  }
  description?: string
}

export function MetricCard({ title, value, change, description }: MetricCardProps) {
  return (
    <Card className="bg-pastel-charcoal border-pastel-soft-gray">
      <CardHeader className="pb-2">
        <CardTitle className="text-sm font-medium text-pastel-soft-white">
          {title}
        </CardTitle>
      </CardHeader>
      <CardContent>
        <div className="text-2xl font-bold text-pastel-soft-white">
          {value}
        </div>
        {change && (
          <div className={`text-xs ${
            change.type === 'increase' 
              ? 'text-pastel-mint-green' 
              : 'text-pastel-blush-pink'
          }`}>
            {change.type === 'increase' ? '+' : '-'}{Math.abs(change.value)}%
          </div>
        )}
        {description && (
          <p className="text-xs text-pastel-soft-white/60 mt-1">
            {description}
          </p>
        )}
      </CardContent>
    </Card>
  )
}
```

---

## Animation System

### anime.js Integration

#### Base Animation Configuration
```javascript
// animations/config.js
export const ANIMATION_CONFIG = {
  // Easing functions
  easing: {
    smooth: 'easeInOutQuad',
    bounce: 'easeOutBounce',
    elastic: 'easeOutElastic',
    cyber: 'easeInOutCubic',
  },
  
  // Duration presets
  duration: {
    fast: 200,
    medium: 400,
    slow: 600,
    ultra: 1000,
  },
  
  // Delay presets
  delay: {
    none: 0,
    short: 100,
    medium: 200,
    long: 400,
  }
}
```

#### Cyber Theme Animations
```javascript
// animations/cyber.js
import anime from 'animejs'

export const cyberAnimations = {
  // Glitch effect for buttons
  glitch: (selector) => {
    anime({
      targets: selector,
      translateX: [
        { value: -2, duration: 50 },
        { value: 2, duration: 50 },
        { value: 0, duration: 50 }
      ],
      boxShadow: [
        { value: '2px 0 var(--cyber-neon-pink)', duration: 50 },
        { value: '-2px 0 var(--cyber-electric-cyan)', duration: 50 },
        { value: '0 0 transparent', duration: 50 }
      ],
      easing: 'linear',
      loop: 2
    })
  },
  
  // Matrix-style text reveal
  matrixReveal: (selector) => {
    anime({
      targets: selector,
      opacity: [0, 1],
      translateY: [20, 0],
      color: [
        { value: 'var(--cyber-matrix-green)', duration: 100 },
        { value: 'var(--cyber-pure-white)', duration: 300 }
      ],
      easing: 'easeOutQuad',
      duration: 600,
      delay: anime.stagger(50)
    })
  },
  
  // Cyber pulse effect
  pulse: (selector) => {
    anime({
      targets: selector,
      scale: [1, 1.05, 1],
      boxShadow: [
        { value: '0 0 0 var(--cyber-electric-cyan)', duration: 300 },
        { value: '0 0 20px var(--cyber-electric-cyan)', duration: 300 },
        { value: '0 0 0 var(--cyber-electric-cyan)', duration: 300 }
      ],
      easing: 'easeInOutQuad',
      duration: 900,
      loop: true
    })
  }
}
```

#### Pastel Theme Animations
```javascript
// animations/pastel.js
import anime from 'animejs'

export const pastelAnimations = {
  // Gentle fade in
  fadeIn: (selector) => {
    anime({
      targets: selector,
      opacity: [0, 1],
      translateY: [10, 0],
      easing: 'easeOutQuad',
      duration: 400,
      delay: anime.stagger(100)
    })
  },
  
  // Soft hover effect
  softHover: (selector) => {
    anime({
      targets: selector,
      scale: [1, 1.02],
      boxShadow: [
        '0 4px 20px rgba(179, 229, 252, 0.1)',
        '0 8px 30px rgba(179, 229, 252, 0.2)'
      ],
      easing: 'easeOutQuad',
      duration: 300
    })
  },
  
  // Data visualization entrance
  chartEnter: (selector) => {
    anime({
      targets: selector,
      opacity: [0, 1],
      scale: [0.8, 1],
      easing: 'easeOutElastic(1, .6)',
      duration: 800,
      delay: anime.stagger(150)
    })
  }
}
```

#### Component-Specific Animations
```typescript
// hooks/useAnimations.ts
import { useEffect } from 'react'
import { cyberAnimations, pastelAnimations } from '@/animations'

export function useCyberAnimations() {
  useEffect(() => {
    // Auto-animate cyber elements
    cyberAnimations.matrixReveal('.cyber-text')
    cyberAnimations.pulse('.cyber-status')
  }, [])
  
  return {
    triggerGlitch: (selector: string) => cyberAnimations.glitch(selector),
    triggerPulse: (selector: string) => cyberAnimations.pulse(selector),
  }
}

export function usePastelAnimations() {
  useEffect(() => {
    // Auto-animate pastel elements
    pastelAnimations.fadeIn('.pastel-card')
    pastelAnimations.chartEnter('.chart-container')
  }, [])
  
  return {
    triggerSoftHover: (selector: string) => pastelAnimations.softHover(selector),
    triggerChartEnter: (selector: string) => pastelAnimations.chartEnter(selector),
  }
}
```

---

## Implementation Guide

### Project Structure
```
src/
├── styles/
│   ├── globals.css           # Global styles and CSS variables
│   ├── themes/
│   │   ├── cyber.css         # Cyber theme overrides
│   │   └── pastel.css        # Pastel theme overrides
│   └── components/
│       ├── ui/               # shadcn/ui components
│       └── custom/           # Custom components
├── animations/
│   ├── config.js            # Animation configuration
│   ├── cyber.js             # Cyber theme animations
│   └── pastel.js            # Pastel theme animations
├── components/
│   ├── ui/                  # shadcn/ui components
│   ├── layout/              # Layout components
│   └── features/            # Feature-specific components
└── hooks/
    ├── useAnimations.ts     # Animation hooks
    └── useTheme.ts          # Theme management
```

### Global Styles Setup
```css
/* styles/globals.css */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Space+Mono:wght@400;700&display=swap');

/* CSS Variables */
:root {
  /* All color and typography variables defined above */
}

/* Base Styles */
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

html {
  font-family: var(--font-ui);
  font-size: 16px;
  line-height: 1.5;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

body {
  background: var(--surface-primary);
  color: var(--text-primary);
  font-size: var(--text-base);
  line-height: var(--leading-normal);
}

/* Scrollbar Styling */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: var(--surface-secondary);
}

::-webkit-scrollbar-thumb {
  background: var(--neutral-600);
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: var(--neutral-500);
}
```

### Theme Provider Setup
```typescript
// providers/ThemeProvider.tsx
'use client'

import { createContext, useContext, useEffect, useState } from 'react'

type Theme = 'cyber' | 'pastel' | 'system'

interface ThemeContextType {
  theme: Theme
  setTheme: (theme: Theme) => void
}

const ThemeContext = createContext<ThemeContextType | undefined>(undefined)

export function ThemeProvider({ children }: { children: React.ReactNode }) {
  const [theme, setTheme] = useState<Theme>('system')
  
  useEffect(() => {
    const root = window.document.documentElement
    root.classList.remove('cyber', 'pastel')
    
    if (theme === 'cyber') {
      root.classList.add('cyber')
    } else if (theme === 'pastel') {
      root.classList.add('pastel')
    }
  }, [theme])
  
  return (
    <ThemeContext.Provider value={{ theme, setTheme }}>
      {children}
    </ThemeContext.Provider>
  )
}

export function useTheme() {
  const context = useContext(ThemeContext)
  if (!context) {
    throw new Error('useTheme must be used within a ThemeProvider')
  }
  return context
}
```

### Application-Specific Implementations

#### NextJS Frontend Example
```typescript
// app/layout.tsx
import { ThemeProvider } from '@/providers/ThemeProvider'
import { Inter } from 'next/font/google'
import '@/styles/globals.css'
import '@/styles/themes/cyber.css'

const inter = Inter({ subsets: ['latin'] })

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" className={inter.className}>
      <body>
        <ThemeProvider>
          {children}
        </ThemeProvider>
      </body>
    </html>
  )
}
```

#### Python Panel Dashboard Example
```python
# dashboard.py
import panel as pn
from pathlib import Path

# Load CSS
css_path = Path(__file__).parent / "styles" / "dashboard.css"
pn.config.raw_css = css_path.read_text()

# Configure Panel
pn.extension('tabulator', 'bokeh', template='material')

# Custom styling
DASHBOARD_CSS = """
.bk-root {
  --primary-color: var(--pastel-sky-blue);
  --secondary-color: var(--pastel-lavender);
  --success-color: var(--pastel-mint-green);
  --error-color: var(--pastel-blush-pink);
  --warning-color: var(--pastel-cream-yellow);
  
  font-family: var(--font-ui);
  background: var(--pastel-midnight-black);
  color: var(--pastel-soft-white);
}

.panel-widget-box {
  background: var(--pastel-charcoal);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  padding: 1rem;
  margin-bottom: 1rem;
}
"""

pn.config.raw_css = DASHBOARD_CSS

# Create dashboard
def create_dashboard():
    return pn.template.MaterialTemplate(
        title="Analytics Dashboard",
        sidebar=[
            pn.pane.Markdown("## Navigation", css_classes=['panel-widget-box']),
            # ... sidebar content
        ],
        main=[
            pn.pane.Markdown("# Dashboard", css_classes=['panel-widget-box']),
            # ... main content
        ],
        header_background=DASHBOARD_CSS,
    )
```

#### CLI Application Example
```python
# cli.py
import click
from rich.console import Console
from rich.theme import Theme

# Custom theme using design system colors
CLI_THEME = Theme({
    "success": "bold green",      # Matrix Green
    "error": "bold red",          # Neon Pink
    "warning": "bold yellow",     # Laser Yellow
    "info": "bold cyan",          # Electric Cyan
    "primary": "bold white",      # Pure White
    "secondary": "dim white",     # Muted white
    "accent": "bold cyan",        # Electric Cyan
})

console = Console(theme=CLI_THEME)

@click.command()
@click.option('--verbose', '-v', is_flag=True, help='Enable verbose output')
def main(verbose):
    """CLI application with cyber theme."""
    console.print("🚀 System initialized", style="success")
    console.print("⚡ Loading modules...", style="info")
    
    if verbose:
        console.print("Debug mode enabled", style="warning")
    
    # ... application logic
```

### Performance Optimization

#### Font Loading Strategy
```css
/* Preload critical fonts */
@font-face {
  font-family: 'Inter';
  src: url('./fonts/Inter-Regular.woff2') format('woff2');
  font-weight: 400;
  font-style: normal;
  font-display: swap;
}

@font-face {
  font-family: 'Space Mono';
  src: url('./fonts/SpaceMono-Regular.woff2') format('woff2');
  font-weight: 400;
  font-style: normal;
  font-display: swap;
}
```

#### CSS Optimization
```css
/* Critical CSS inlining */
.critical-above-fold {
  font-family: var(--font-ui);
  background: var(--surface-primary);
  color: var(--text-primary);
}

/* Non-critical CSS lazy loading */
@media (prefers-reduced-motion: no-preference) {
  /* Animations only when user prefers motion */
  .animate-on-scroll {
    animation: fadeIn 0.6s ease-out;
  }
}
```

---

## Accessibility Guidelines

### Color Contrast Requirements
- **Cyber Theme**: Minimum 7:1 contrast ratio (AAA)
- **Pastel Theme**: Minimum 4.5:1 contrast ratio (AA)
- **Text on colored backgrounds**: Always test contrast
- **Interactive elements**: Minimum 44px touch target

### Focus Management
```css
/* Focus indicators */
.focus-visible {
  outline: 2px solid var(--cyber-electric-cyan);
  outline-offset: 2px;
}

.pastel-focus {
  outline: 2px solid var(--pastel-sky-blue);
  outline-offset: 2px;
}
```

### Motion Preferences
```css
@media (prefers-reduced-motion: reduce) {
  *,
  *::before,
  *::after {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
    scroll-behavior: auto !important;
  }
}
```

---

## Testing & Quality Assurance

### Visual Testing Checklist
- [ ] Color contrast meets accessibility standards
- [ ] Typography scales properly across devices
- [ ] Interactive elements have clear focus states
- [ ] Animations respect motion preferences
- [ ] Themes switch properly without layout shifts
- [ ] Components render correctly in both themes

### Performance Testing
- [ ] Fonts load efficiently with proper fallbacks
- [ ] CSS bundle size is optimized
- [ ] Animations don't cause layout thrashing
- [ ] Theme switching is smooth and fast

### Cross-platform Testing
- [ ] NextJS frontend works across browsers
- [ ] Python Panel dashboard renders correctly
- [ ] CLI colors display properly in terminals
- [ ] Mobile responsiveness maintained

---

*This `.rules_ui` file provides comprehensive guidelines for implementing a cohesive design system across NextJS frontend, Python Panel dashboard, and CLI applications using the Cyber and Pastel themes with modern typography and animations.*