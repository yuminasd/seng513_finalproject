import Logout from "./logout";
import Image from 'next/image';

export default function Navbar() {
    return (
        <div className="absolute top-0 left-0  w-full bg-black border-b border-neutral-800 flex justify-between items-center p-4 z-50">
            <a className="font-xl text-purple-500 flex  items-center gap-4 font-bold text-lg hover:text-purple-600 " href="/">
                <Image src="/images/logo.svg" alt="Groups" width={36} height={36} />
                Movies Match
            </a>

            <div className="flex gap-4">
                <a className="font-xl text-white rounded bg-white bg-opacity-10 p-2 px-4 hover:text-gray-300" href="/profile">Profile</a>
                <Logout />
            </div>
        </div>
    )
}