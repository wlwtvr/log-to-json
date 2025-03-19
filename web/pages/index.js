import Head from 'next/head';
import ParserComponent from '../components/ParserComponent';

export default function Home() {
  return (
    <div>
      <Head>
        <title>Log2JSON Parser</title>
        <meta name="description" content="Parse logs to JSON format" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className="flex flex-col items-center justify-center min-h-screen py-2 bg-gray-50">
        <h1 className="text-4xl font-bold mb-4">Log2JSON Parser</h1>
        <ParserComponent />
      </main>
    </div>
  );
}
