'use client'
import { useRouter } from 'next/navigation'

export default function Page() {
    const router = useRouter();
    // Your login logic here...
    const handleLogin = () => {
        // Perform login actions...
        router.push('/');
    };

    return (
        <div>
            <input placeholder="email" />
            <input placeholder="password" />
            <button onClick={handleLogin}>Login</button>
        </div>
    );
}