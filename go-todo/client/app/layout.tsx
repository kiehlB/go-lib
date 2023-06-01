import Link from "next/link";
import "../styles/globals.css";
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "todo-app",
  description: "todo list",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang='ko'>
      <body className={inter.className}>
        <Link
          href='/'
          className='mx-auto grid max-w-[98.5rem] mxl:max-w-[75rem] mmd:grid-cols-10 text-3xl py-4 font-semibold'
        >
          Todo App
        </Link>
        {children}
      </body>
    </html>
  );
}
