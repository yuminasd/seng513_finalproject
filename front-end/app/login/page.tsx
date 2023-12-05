'use client'
import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Button from '../components/button';

export default function Page() {
    const router = useRouter();
    const [showPassword, setShowPassword] = useState(false);
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleLogin = async () => {
        try {
            const response = await fetch('http://localhost:5000/checklogin', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ "email":email, "password":password }),
            });
            console.log(response)
            if (response.ok) {
                // Login successful, redirect to the desired page
                router.push('/?userid=656500413c49f1af1a59b5d1');
                //get user id from response and append to string
            } else {
                // Handle login failure****
                console.error('Login failed');
            }
        } catch (error) {
            // Handle fetch error
            console.error('Error during login:', error);
        }
    };

    return (
        <div className="flex items-center justify-center h-screen bg-gray-900 text-white">
            <div className="w-120 text-center">
                {/* MovieMatch Logo */}
                <img src="/logo.png" alt="logo" className="mb-4" />

                <h1 className="text-3xl font-bold mb-4">
                    <span className="text-white">Welcome To</span>
                    <span className="text-purple-500"> Movie Match</span>
                </h1>
                <div className="mb-4 w-120">
                    <label className="text-white block mb-1 text-left">Email</label>
                    <input
                        className="my-2 p-2 w-full border rounded text-gray-500 bg-gray-800 border-transparent"
                        type="text"
                        placeholder="Email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                </div>

                <div className="mb-4 w-120 relative">
                    <label className="text-white block mb-1 text-left">Password</label>
                    <div className="flex items-center">
                        <input
                            className="my-2 p-2 w-full border rounded text-gray-500 bg-gray-800 border-transparent"
                            type={showPassword ? 'text' : 'password'}
                            placeholder="Password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                        />
                        <span
                            className="text-gray-400 text-sm cursor-pointer ml-2 bg-gray-800 p-2 rounded"
                            onClick={() => setShowPassword(!showPassword)}
                        >
                            {showPassword ? '👁️' : '👁️'}
                        </span>
                    </div>
                </div>

                <Button text="Sign In" color="primary" onClick={handleLogin} />
                <div className="mt-2 text-gray-400 text-sm cursor-pointer">Create Account</div>
            </div>
        </div>
    );
}
