'use client'
// login/page.tsx

import { useState } from 'react';
import { useRouter } from 'next/router';
import Link from 'next/link';
import Button from '../components/button';

export default function Page() {
  const router = useRouter();
  const [showPassword, setShowPassword] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [showError, setShowError] = useState(false);

  const handleLogin = async () => {
    try {
      const response = await fetch('http://localhost:5000/checklogin', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email: email, password: password }),
      });

      if (response.ok) {
        try {
          const responseData = await response.json();
          console.log(responseData);

          const userId = responseData.data.id;
          const userRole = responseData.data.userRole;

          localStorage.setItem('userId', userId);
          localStorage.setItem('userRole', userRole);

          router.push(`/?userid=${userId}`);
        } catch (error) {
          console.error('Error parsing JSON in the API response:', error);
        }
      } else {
        console.error('Login failed');
        setShowError(true);
      }
    } catch (error) {
      console.error('Error during login:', error);
    }
  };

  const storedUserId = localStorage.getItem('userId');

  return (
    <div className="flex items-center justify-center h-screen bg-gray-900 text-white">
      <div className="w-120 text-center">
        <div className="mb-4">
          <img src="/logo.png" alt="logo" className="mx-auto" />
        </div>

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
        <Link href="/app/signup/page">
          <a className="mt-2 text-gray-400 text-sm cursor-pointer">Create Account</a>
        </Link>

        {showError && (
          <div className="absolute top-0 left-0 right-0 bottom-0 flex items-center justify-center">
            <div className="bg-gray-800 p-4 rounded shadow-md text-white">
              <p className="text-red-500">Invalid email or password. Please try again.</p>
              <button className="mt-2 p-2 bg-purple-500 text-white rounded" onClick={() => setShowError(false)}>
                Close
              </button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
