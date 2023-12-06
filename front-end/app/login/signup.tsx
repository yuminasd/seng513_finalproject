// login/signup.tsx

import { useState } from 'react';
import { useRouter } from 'next/router';

export default function SignUp() {
  const router = useRouter();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSignUp = async () => {
    try {
      const response = await fetch('http://localhost:5000/users', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: email,
          password: password,
          // Add any additional fields needed for user creation
        }),
      });

      if (response.ok) {
        // Redirect to the login page after successful signup
        router.push('/login/page');
      } else {
        console.error('Account creation failed');
      }
    } catch (error) {
      console.error('Error during account creation:', error);
    }
  };

  return (
    <div className="flex items-center justify-center h-screen bg-gray-900 text-white">
      <div className="w-120 text-center">
        <h1 className="text-3xl font-bold mb-4">
          <span className="text-white">Create Your</span>
          <span className="text-purple-500"> Account</span>
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
          <input
            className="my-2 p-2 w-full border rounded text-gray-500 bg-gray-800 border-transparent"
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>

        {/* Button component included here */}
        <button
          className="my-2 p-2 w-full bg-purple-500 text-white rounded cursor-pointer"
          onClick={handleSignUp}
        >
          Sign Up
        </button>
      </div>
    </div>
  );
}
