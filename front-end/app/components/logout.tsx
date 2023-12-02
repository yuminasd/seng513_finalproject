'use client'
import { useRouter } from 'next/navigation'
const Logout = () => {
    const router = useRouter()

    // Your logout logic here...
    const handleLogout = () => {
        // Perform logout actions...

        // Redirect to the login page after logout
        router.push('/login');
    };

    return (
        <div>
            {/* Your logout button or trigger */}
            <button className="font-xl text-white rounded bg-white bg-opacity-10 p-2 px-4 hover:text-red-500 " onClick={handleLogout}>Logout</button>
        </div>
    );
};

export default Logout;