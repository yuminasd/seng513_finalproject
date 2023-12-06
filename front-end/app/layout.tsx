'use client'
import { Inter } from 'next/font/google'
import './globals.css'


const inter = Inter({ subsets: ['latin'] })


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
