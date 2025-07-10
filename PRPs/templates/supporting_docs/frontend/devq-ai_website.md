// Required design system implementation
// frontend/src/styles/globals.css
// Key Design System Integration:
// CSS Variables - Complete implementation of cyber and pastel color palettes
// Typography System - Inter Nerd Font for UI, Space Mono for code/terminal
// Theme Switcher - Dynamic switching between cyber and pastel themes
// Consistent Styling - Utility classes for both color schemes
// Animated Components - Framer Motion with design system colors
// Terminal Aesthetics - Cyber-themed terminal with proper status colors
// Accessibility - Proper contrast ratios and font sizing
// Component Library - Reusable components following design system patterns

@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Space+Mono:wght@400;700&display=swap');

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

  /* Pastel Black Palette */
  --pastel-midnight-black: #000000;
  --pastel-charcoal: #0f0f0f;
  --pastel-soft-gray: #1e1e1e;
  --pastel-soft-white: #f8f8f8;
  --pastel-mint-green: #a8e6a3;
  --pastel-blush-pink: #ffb3ba;
  --pastel-sky-blue: #b3e5fc;
  --pastel-cream-yellow: #fff9c4;

  /* Typography */
  --font-ui: 'Inter', 'Inter Nerd Font', 'Segoe UI', 'Roboto', sans-serif;
  --font-mono: 'Space Mono', 'Space Mono Nerd Font', 'JetBrains Mono', monospace;
  --font-light: 300;
  --font-regular: 400;
  --font-medium: 500;
  --font-semibold: 600;
  --font-bold: 700;

  /* Font Sizes */
  --text-xs: 0.75rem;
  --text-sm: 0.875rem;
  --text-base: 1rem;
  --text-lg: 1.125rem;
  --text-xl: 1.25rem;
  --text-2xl: 1.5rem;
  --text-3xl: 1.875rem;
  --text-4xl: 2.25rem;
}

/* Utility Classes */
.font-ui { font-family: var(--font-ui); }
.font-mono { font-family: var(--font-mono); }

/* Cyber Theme Classes */
.cyber-bg-primary { background-color: var(--cyber-void-black); }
.cyber-bg-secondary { background-color: var(--cyber-carbon-black); }
.cyber-bg-surface { background-color: var(--cyber-gray); }
.cyber-text-primary { color: var(--cyber-white); }
.cyber-text-success { color: var(--cyber-matrix-green); }
.cyber-text-error { color: var(--cyber-neon-pink); }
.cyber-text-warning { color: var(--cyber-laser-yellow); }
.cyber-text-info { color: var(--cyber-electric-cyan); }

/* Pastel Theme Classes */
.pastel-bg-primary { background-color: var(--pastel-midnight-black); }
.pastel-bg-secondary { background-color: var(--pastel-charcoal); }
.pastel-bg-surface { background-color: var(--pastel-soft-gray); }
.pastel-text-primary { color: var(--pastel-soft-white); }
.pastel-text-success { color: var(--pastel-mint-green); }
.pastel-text-error { color: var(--pastel-blush-pink); }
.pastel-text-warning { color: var(--pastel-cream-yellow); }
.pastel-text-info { color: var(--pastel-sky-blue); }

// Required NextJS app structure with design system
// frontend/src/app/layout.tsx
import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: 'DevQ.ai - AI-Powered Development Platform',
  description: 'Intelligent agents for modern development workflows',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className="font-ui cyber-bg-primary cyber-text-primary">
        {children}
      </body>
    </html>
  )
}

// Required hero section with cyber theme
// frontend/src/components/HeroSection.tsx
'use client'

import { motion } from 'framer-motion'
import { useState, useEffect } from 'react'

export function HeroSection() {
  const [currentAgent, setCurrentAgent] = useState(0)
  const agents = ['superagent', 'codeagent', 'datascienceagent', 'cloudagent']

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentAgent((prev) => (prev + 1) % agents.length)
    }, 3000)
    return () => clearInterval(interval)
  }, [])

  return (
    <section className="min-h-screen cyber-bg-primary flex items-center justify-center">
      <div className="container mx-auto px-4 py-20">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8 }}
          className="text-center"
        >
          <h1 className="font-ui font-bold cyber-text-primary mb-6" 
              style={{ fontSize: 'var(--text-4xl)' }}>
            Meet Your AI Development Team
          </h1>
          <motion.p
            key={currentAgent}
            initial={{ opacity: 0, scale: 0.8 }}
            animate={{ opacity: 1, scale: 1 }}
            transition={{ duration: 0.5 }}
            className="font-ui cyber-text-primary mb-8"
            style={{ fontSize: 'var(--text-2xl)' }}
          >
            Currently featuring: {' '}
            <span className="cyber-text-info font-mono font-bold">
              {agents[currentAgent]}
            </span>
          </motion.p>
        </motion.div>
        
        <CyberTerminalAnimation />
      </div>
    </section>
  )
}

// Required cyber terminal animation with design system
// frontend/src/components/CyberTerminalAnimation.tsx
'use client'

import { motion } from 'framer-motion'
import { useState, useEffect } from 'react'

const commands = [
  { 
    command: 'devqai agentical run --workflow data-analysis',
    result: '✓ Analysis complete',
    status: 'success'
  },
  { 
    command: 'devqai ptolemies search --query "API patterns"',
    result: '⚠ 847 results found',
    status: 'warning'
  },
  { 
    command: 'devqai machina list --category development',
    result: '→ Scanning registry...',
    status: 'info'
  },
  { 
    command: 'devqai trustlis auth --generate-proof',
    result: '✗ Permission denied',
    status: 'error'
  }
]

export function CyberTerminalAnimation() {
  const [currentCommand, setCurrentCommand] = useState(0)
  const [displayedText, setDisplayedText] = useState('')
  const [showResult, setShowResult] = useState(false)

  useEffect(() => {
    const command = commands[currentCommand]
    let index = 0
    
    setShowResult(false)
    
    const typeInterval = setInterval(() => {
      if (index < command.command.length) {
        setDisplayedText(command.command.slice(0, index + 1))
        index++
      } else {
        clearInterval(typeInterval)
        setShowResult(true)
        
        setTimeout(() => {
          setCurrentCommand((prev) => (prev + 1) % commands.length)
          setDisplayedText('')
        }, 2000)
      }
    }, 50)

    return () => clearInterval(typeInterval)
  }, [currentCommand])

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'success': return 'cyber-text-success'
      case 'warning': return 'cyber-text-warning' 
      case 'error': return 'cyber-text-error'
      case 'info': return 'cyber-text-info'
      default: return 'cyber-text-primary'
    }
  }

  return (
    <motion.div
      initial={{ opacity: 0, y: 40 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.8, delay: 0.4 }}
      className="cyber-bg-secondary rounded-lg p-6 font-mono max-w-4xl mx-auto"
      style={{ fontSize: 'var(--text-sm)' }}
    >
      <div className="flex items-center mb-4">
        <div className="w-3 h-3 bg-red-500 rounded-full mr-2"></div>
        <div className="w-3 h-3 bg-yellow-500 rounded-full mr-2"></div>
        <div className="w-3 h-3 bg-green-500 rounded-full mr-2"></div>
        <span className="cyber-text-info font-bold ml-4">DEVQAI TERMINAL</span>
      </div>
      
      <div className="space-y-2">
        <div className="flex items-center">
          <span className="cyber-text-info mr-2">$</span>
          <span className="cyber-text-primary">{displayedText}</span>
          <motion.span
            animate={{ opacity: [0, 1, 0] }}
            transition={{ duration: 1, repeat: Infinity }}
            className="cyber-text-info ml-1"
          >
            |
          </motion.span>
        </div>
        
        {showResult && (
          <motion.div
            initial={{ opacity: 0, x: 20 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ duration: 0.3 }}
            className={`ml-4 ${getStatusColor(commands[currentCommand].status)}`}
          >
            {commands[currentCommand].result}
          </motion.div>
        )}
      </div>
    </motion.div>
  )
}

// Required features showcase with pastel theme
// frontend/src/components/FeatureCard.tsx
'use client'

import { motion } from 'framer-motion'

interface FeatureCardProps {
  title: string
  description: string
  icon: string
  index: number
  theme: 'cyber' | 'pastel'
}

export function FeatureCard({ title, description, icon, index, theme }: FeatureCardProps) {
  const isPassive = theme === 'pastel'
  
  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5, delay: index * 0.1 }}
      whileHover={{ scale: 1.05 }}
      className={`
        ${isPassive ? 'pastel-bg-surface' : 'cyber-bg-surface'} 
        rounded-lg p-6 transition-all duration-300
        hover:shadow-lg border border-opacity-20
        ${isPassive ? 'border-white' : 'border-white'}
      `}
    >
      <div className="text-4xl mb-4">{icon}</div>
      <h3 className={`
        font-ui font-semibold mb-3
        ${isPassive ? 'pastel-text-primary' : 'cyber-text-primary'}
      `} style={{ fontSize: 'var(--text-xl)' }}>
        {title}
      </h3>
      <p className={`
        font-ui font-regular
        ${isPassive ? 'pastel-text-primary' : 'cyber-text-primary'}
      `} style={{ fontSize: 'var(--text-base)' }}>
        {description}
      </p>
      
      {/* Theme-specific accent */}
      <div className={`
        mt-4 h-1 w-16 rounded-full
        ${getAccentColor(title, theme)}
      `} />
    </motion.div>
  )
}

function getAccentColor(title: string, theme: 'cyber' | 'pastel') {
  const colorMap = {
    cyber: {
      'Agentical': 'bg-green-500', // Matrix green
      'Ptolemies': 'bg-cyan-500',  // Electric cyan
      'Machina': 'bg-yellow-500',  // Laser yellow
      'Trustlis': 'bg-pink-500'    // Neon pink
    },
    pastel: {
      'Agentical': 'bg-green-200', // Mint green
      'Ptolemies': 'bg-blue-200',  // Sky blue
      'Machina': 'bg-yellow-200',  // Cream yellow
      'Trustlis': 'bg-pink-200'    // Blush pink
    }
  }
  
  return colorMap[theme][title] || 'bg-gray-500'
}

// Required features page with theme switcher
// frontend/src/app/features/page.tsx
'use client'

import { useState } from 'react'
import { FeatureCard } from '@/components/FeatureCard'
import { motion } from 'framer-motion'

const features = [
  {
    title: 'Agentical',
    description: 'Intelligent agents for automated workflows using PydanticAI',
    icon: '🤖'
  },
  {
    title: 'Ptolemies',
    description: 'Knowledge management with RAG and graph databases',
    icon: '📚'
  },
  {
    title: 'Machina',
    description: 'MCP registry for extensible tool integration',
    icon: '🔧'
  },
  {
    title: 'Trustlis',
    description: 'Zero-knowledge authentication with zkSTARK',
    icon: '🔒'
  }
]

export default function FeaturesPage() {
  const [theme, setTheme] = useState<'cyber' | 'pastel'>('cyber')

  return (
    <div className={`
      min-h-screen py-20 transition-all duration-500
      ${theme === 'cyber' ? 'cyber-bg-primary' : 'pastel-bg-primary'}
    `}>
      <div className="container mx-auto px-4">
        {/* Theme Switcher */}
        <div className="flex justify-center mb-8">
          <div className={`
            ${theme === 'cyber' ? 'cyber-bg-secondary' : 'pastel-bg-secondary'}
            rounded-full p-1 flex gap-1
          `}>
            <button
              onClick={() => setTheme('cyber')}
              className={`
                px-4 py-2 rounded-full font-ui font-medium transition-all
                ${theme === 'cyber' 
                  ? 'cyber-bg-surface cyber-text-info' 
                  : 'cyber-text-primary hover:cyber-bg-surface'
                }
              `}
            >
              CYBER
            </button>
            <button
              onClick={() => setTheme('pastel')}
              className={`
                px-4 py-2 rounded-full font-ui font-medium transition-all
                ${theme === 'pastel' 
                  ? 'pastel-bg-surface pastel-text-info' 
                  : 'pastel-text-primary hover:pastel-bg-surface'
                }
              `}
            >
              PASTEL
            </button>
          </div>
        </div>

        <motion.h1 
          key={theme}
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          className={`
            font-ui font-bold text-center mb-16
            ${theme === 'cyber' ? 'cyber-text-primary' : 'pastel-text-primary'}
          `}
          style={{ fontSize: 'var(--text-4xl)' }}
        >
          {theme === 'cyber' ? 'SYSTEM_MODULES.EXE' : 'Development Tools'}
        </motion.h1>

        <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
          {features.map((feature, index) => (
            <FeatureCard
              key={`${feature.title}-${theme}`}
              {...feature}
              index={index}
              theme={theme}
            />
          ))}
        </div>
      </div>
    </div>
  )
}

// Required package.json with design system dependencies
// frontend/package.json
{
  "name": "devqai-frontend",
  "version": "0.1.0",
  "private": true,
  "scripts": {
    "dev": "next dev",
    "build": "next build", 
    "start": "next start",
    "lint": "next lint"
  },
  "dependencies": {
    "next": "14.0.0",
    "react": "^18",
    "react-dom": "^18",
    "framer-motion": "^10.16.5",
    "typescript": "^5"
  },
  "devDependencies": {
    "@types/node": "^20",
    "@types/react": "^18", 
    "@types/react-dom": "^18",
    "eslint": "^8",
    "eslint-config-next": "14.0.0"
  }
}