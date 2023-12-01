'use client'
// Import necessary components
import { useRouter } from 'next/navigation';
import Button from '../components/button';

export default function Page() {
    const router = useRouter();

    // Your login logic here...
    const handleLogin = () => {
        // Perform login actions...
        router.push('/');
    };

    return (
        <div className="flex items-center justify-center h-screen">
            <div className="w-64 text-center">
                <h1 className="text-3xl font-bold mb-4 text-purple-500">Please Sign In</h1>
                <input className="my-2 p-2 w-full border border-gray-300 rounded text-black" type="text" placeholder="Email" />
                <input className="my-2 p-2 w-full border border-gray-300 rounded text-black" type="password" placeholder="Password" />
                <Button text="Login" color="primary" onClick={handleLogin} />
            </div>
        </div>
    );
    }