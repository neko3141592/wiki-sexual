"use client";
import { useState } from "react";
import ArticleSearchInput from "@/components/ArticleSearchInput";

export default function Home() {
    const [start, setStart] = useState("");
    const [end, setEnd] = useState("");
    const [loading, setLoading] = useState(false);
    const [path, setPath] = useState<string[]>([]);
    const [error, setError] = useState("");

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true);
        setError("");
        setPath([]);
        try {
            const res = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/api/path?start=${encodeURIComponent(start)}&end=${encodeURIComponent(end)}`
            );
            const data = await res.json();
            if (!res.ok) {
                setError(data.error || data.message || "エラーが発生しました");
            } else {
                setPath(data.path || []);
            }
        } catch {
            setError("通信エラー");
        } finally {
            setLoading(false);
        }
    };

    return (
        <>
            <style jsx global>{`
                @keyframes spin-slow {
                    from { transform: rotate(0deg); }
                    to { transform: rotate(360deg); }
                }
                @keyframes spin-reverse-slow {
                    from { transform: rotate(0deg); }
                    to { transform: rotate(-360deg); }
                }
                .animate-spin-slow {
                    animation: spin-slow 4s linear infinite;
                }
                .animate-spin-reverse-slow {
                    animation: spin-reverse-slow 4s linear infinite;
                }
            `}</style>
            <div className="min-h-screen bg-zinc-900 text-white flex items-center justify-center p-8 relative">
            {loading && (
                    <div className="fixed inset-0 bg-black/70 z-50 flex items-center justify-center">
                        <div className="flex flex-col items-center">
                            <div className="relative w-20 h-20">
                                <div className="absolute inset-0 border-4 border-t-4 border-t-[#00ADD8] border-transparent rounded-full animate-spin-slow"></div>
                                <div className="absolute inset-2 border-4 border-b-4 border-b-[#00ADD8] border-transparent rounded-full animate-spin-reverse-slow"></div>
                                <div className="absolute inset-4 border-4 border-r-4 border-r-[#00ADD8] border-transparent rounded-full animate-spin-slow"></div>
                                <div className="absolute inset-6 border-4 border-l-4 border-l-[#00ADD8] border-transparent rounded-full animate-spin-reverse-slow"></div>
                            </div>
                            <p className="mt-6 text-xl font-medium text-[#00ADD8] animate-pulse">
                                探索中...
                            </p>
                        </div>
                    </div>
            )}
            <div className="container max-w-xl w-full">
                <h1 className="text-5xl font-bold mb-6 tracking-tight text-center">
                    Wiki<span className="text-[#00ADD8] italic">GO</span>lfer
                </h1>
                {error && (
                    <div className="flex items-center justify-center mb-8">
                        <div className="flex items-center gap-3 px-6 py-4 bg-red-500/10 border border-red-500/30 rounded-lg backdrop-blur-sm">
                            <svg className="w-6 h-6 text-red-400 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                            </svg>
                            <span className="text-red-400 text-lg font-medium">{error}</span>
                        </div>
                    </div>
                )}
                <form
                    onSubmit={handleSubmit}
                    className="flex flex-col gap-4 mb-12"
                >
                    <ArticleSearchInput
                        value={start}
                        onChange={setStart}
                        placeholder="開始記事タイトルを入力"
                    />
                    <ArticleSearchInput
                        value={end}
                        onChange={setEnd}
                        placeholder="終点記事タイトルを入力"
                    />
                    <button
                        type="submit"
                        disabled={loading}
                        className="px-4 py-4 font-bold bg-white text-black rounded-lg cursor-pointer hover:bg-[#00ADD8] hover:text-white transition-colors duration-300 disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:bg-white disabled:hover:text-black"
                    >
                        {loading ? "検索中..." : "検索"}
                    </button>
                </form>
                <div className="p-4 bg-gradient-to-br from-[#0a0a0a] to-[#1a1a1a] border-2 border-[#333] rounded-xl min-h-[200px] text-base leading-relaxed">
                    {path.length ? (
                        <div className="space-y-2">
                            {path.map((title, index) => (
                                <div key={index}>
                                    <div className="flex items-center gap-4 py-3 px-2 rounded-lg hover:bg-white/5 transition-all duration-200 group">
                                        <span className="text-[#00ADD8] font-bold min-w-[2.5rem] text-right text-lg">
                                            {index + 1}
                                        </span>
                                        <a
                                            href={`https://ja.wikipedia.org/wiki/${encodeURIComponent(
                                                title
                                            )}`}
                                            target="_blank"
                                            rel="noopener noreferrer"
                                            className="text-white  flex-1"
                                        >
                                            {title}
                                        </a>
                                    </div>
                                </div>
                            ))}
                        </div>
                    ) : error ? (
                        <div className="flex items-center justify-center h-32 text-red-400">
                            検索結果がありません。
                        </div>
                    ) : (
                        <div className="flex items-center justify-center h-32 text-gray-500">
                            結果がここに表示されます
                        </div>
                    )}
                </div>
            </div>
        </div>
        <div className="fixed bottom-4 right-4 text-gray-500 text-sm">
            © 2025 WikiGOlfer
        </div>
        </>
    );
}
