import type { Metadata } from 'next';

import { Inter } from 'next/font/google';
import './globals.css';
import QueryProvider from './provider/QueryProvider';
import AuthProvider from './provider/AuthProvider';
import Layout from './components/Layout';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'BondsApp',
  description: 'Generated by create next app',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang='en'>
      <QueryProvider>
        <AuthProvider>
          <body className={inter.className}>
            <Layout>
              {children}
            </Layout>
          </body>
        </AuthProvider>
      </QueryProvider>
    </html>
  );
}
