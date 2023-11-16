import Logout from "./logout";

export default function Navbar() {
    return (
        <div className="absolute top-0 w-full bg-gray-800 flex justify-between p-4 ">
            <a className="font-xl text-white" href="/">Groups</a>
            <div className="flex gap-4">
                <a className="font-xl text-white" href="/profile">Profile</a>
                <Logout />
            </div>
        </div>
    )
}