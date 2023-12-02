import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Movie Match',
  description: 'Match Movies',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" className="w-screen h-screen pt-20 overflow-hidden flex flex-col">
      <body className={inter.className}>{children}</body>
    </html>
  )
}
