import Navbar from "../components/navbar";

export default function Home() {
    return (
        <main className="flex min-h-screen min-w-screen">
            <Navbar />
            <div className="w-5/12 pt-24 p-4 border-r border-gray-700">
                <div className="w-full flex flex-col gap-2 p-16 border border-gray-700 rounded-2xl">
                    <h1>Group Name</h1>
                    <p> Group Code</p>
                </div>
                <a href="/profile">User </a>
            </div>


            <div className="w-full pt-24 p-4  grid grid-cols-4 justify-center">
                <p className="text-center">Movie Component</p>
                <p className="text-center">Movie Component</p>
                <p className="text-center">Movie Component</p>
                <p className="text-center">Movie Component</p>
                <p className="text-center">Movie Component</p>
                <p className="text-center">Movie Component</p>
                <p className="text-center">Movie Component</p>
                <p className="text-center">Movie Component</p>
                <p className="text-center">Movie Component</p>
                <p className="text-center">Movie Component</p>
                <p className="text-center">Movie Component</p>
                <p className="text-center">Movie Component</p>
            </div>
        </main>
    )
}
