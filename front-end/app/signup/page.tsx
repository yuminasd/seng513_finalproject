'use client'
// Import necessary dependencies
import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Button from '../components/button';

// Create and export the Signup component
export default function Signup() {
    // Initialize necessary state variables
    const router = useRouter();
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [showPassword, setShowPassword] = useState(false);
    const [showError, setShowError] = useState(false);

    // Function to handle signup
    const handleSignup = async () => {
        try {
            const response = await fetch('http://localhost:5000/users', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    name: name,
                    emailAddress: email,
                    password: password,
                    image: '/logo.png', // Hard-coded link to the logo image
                }),
            });

            if (response.ok) {
                // Handle successful signup
                router.push('/login'); // Redirect to the login page
            } else {
                // Handle signup failure
                console.error('Signup failed');
                setShowError(true);
            }
        } catch (error) {
            // Handle fetch error
            console.error('Error during signup:', error);
        }
    };

    return (
        <div className="flex items-center justify-center h-screen bg-gray-900 text-white">
            <div className="w-120 text-center">
                {/* Same logo as the login page */}
                <div className="mb-4">
                    <img src="/logo.png" alt="logo" className="mx-auto" />
                </div>

                <h1 className="text-3xl font-bold mb-4">
                    <span className="text-white">Create Your</span>
                    <span className="text-purple-500"> Movie Match Account</span>
                </h1>
                
                {/* Name input */}
                <div className="mb-4 w-120">
                    <label className="text-white block mb-1 text-left">Name</label>
                    <input
                        className="my-2 p-2 w-full border rounded text-gray-500 bg-gray-800 border-transparent"
                        type="text"
                        placeholder="Name"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                    />
                </div>

                {/* Email input */}
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

                {/* Password input with show/hide toggle */}
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

                {/* Signup button */}
                <Button text="Sign Up" color="primary" onClick={handleSignup} />

                {/* Link to login page */}
                <div className="mt-2 text-gray-400 text-sm cursor-pointer" onClick={() => router.push('/login')}>
                    Already have an account? Login here.
                </div>

                {/* Error pop-up */}
                {showError && (
                    <div className="absolute top-0 left-0 right-0 bottom-0 flex items-center justify-center">
                        <div className="bg-gray-800 p-4 rounded shadow-md text-white">
                            <p className="text-red-500">Signup failed.</p>
                            <button className="mt-2 p-2 bg-purple-500 text-white rounded"
                                onClick={() => setShowError(false)}>
                                Close
                            </button>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
}
