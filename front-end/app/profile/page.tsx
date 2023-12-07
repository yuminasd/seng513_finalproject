'use client';
import { useEffect, useState } from "react";
import Navbar from "../components/navbar";
import { User } from "../types";
import Button from "../components/button";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export default function Profile() {
    const [user, setUser] = useState<User | null>(null);
    const userId = typeof window !== 'undefined' ? localStorage.getItem('userId') : null;
    const userRole = typeof window !== 'undefined' ? localStorage.getItem('userRole') : null;
    const [isReadOnly, setIsReadOnly] = useState(true); // Added state for readOnly

    useEffect(() => {
        const fetchData = async () => {
            try {
                console.log(userRole);
                const response = await fetch('http://localhost:5000/users/' + userId, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                    }
                });
                const data = await response.json();
                // console.log(data.data.data);
                setUser(data.data.data as User);
                // Set readOnly state based on userRole
                setIsReadOnly(userRole !== 'admin');
            } catch (error) {
                console.error('Error fetching groups data:', error);
            }
        };

        fetchData();
    }, []);

    const updateUser = async () => {
        try {
            const response = await fetch('http://localhost:5000/users/' + userId, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify([user]),
            });
        }
        catch (error) {
            toast.error("Fail Update User");
        }
        toast.success("Updated User");
    };


    return (
        <section className=" flex p-4 h-screen">
            <ToastContainer position="bottom-right"
                autoClose={5000}
                z-index={9999}
                hideProgressBar={false}
                newestOnTop={false}
                closeOnClick
                rtl={false}
                pauseOnFocusLoss
                draggable
                pauseOnHover
                theme="dark" />
            <Navbar />
            <form className="flex flex-col gap-4 p-4 w-80 h-full">
                <img className="p-4  w-80 aspect-square border border-neutral-800 rounded" src={user?.image} alt="User Image" />
                <input
                    className="bg-black p-2 border border-neutral-800 rounded"
                    defaultValue={user?.name}
                    readOnly={isReadOnly}
                />
                <input
                    className="bg-black p-2 border border-neutral-800 rounded"
                    defaultValue={user?.id}
                    readOnly={isReadOnly}
                />
            </form>
            {isReadOnly ? (<div></div>) : (<div className="absolute left-0 bottom-0 w-full bg-neutral-900 p-4">
                <Button text={"Update"} color={"primary"} onClick={updateUser} />
            </div>)}


        </section>
    );
}
