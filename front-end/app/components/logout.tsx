'use client'
import { useRouter } from 'next/navigation'
const Logout = () => {
    const router = useRouter()
    const userId = typeof window !== 'undefined' ? localStorage.getItem('userId') : null;
    // Your logout logic here...
    const handleLogout = async () => {

        const response = await fetch('http://localhost:5000/logout/' + userId, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (response.ok) {
            router.push('/login');
        } else {
            // Handle error response
            console.error('Failed to logout:', response.status, response.statusText);
        }
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