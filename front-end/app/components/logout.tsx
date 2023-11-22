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
            <button onClick={handleLogout}>Logout</button>
        </div>
    );
};

export default Logout;