
import React, { useState } from 'react';
import Button from '../button';
import { User } from '@/app/types';

interface ModalProps {
    hidden: boolean;
    onClose: () => void;
    user: User | null;
}

const Modal: React.FC<ModalProps> = ({ hidden, onClose, user }) => {

    const [name, setName] = useState('');
    const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setName(event.target.value);
    };

    const toggleModal = () => {
        onClose();
    };

    const handleJoinClick = async () => {
        try {

            console.log(user);
            const response = await fetch('http://localhost:5000/groups', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    Name: name,
                    Genre: [],
                    members: [user], // Replace 'currentUserId' with the actual user ID
                }),
            });

            if (response.ok) {
                // Handle successful response, e.g., close the modal
                onClose();
            } else {
                // Handle error response
                console.error('Failed to create group:', response.status, response.statusText);
            }
        } catch (error) {
            console.error('Error create group:', error);
        }
    };

    if (!hidden) {
        return (
            <div className="fixed top-0 left-0 w-full h-full">
                <div className="absolute w-full h-full bg-black bg-opacity-50 flex justify-center items-center " >
                    <form className="bg-neutral-900 p-4 text-white flex flex-col z-50 rounded-xl">

                        <span className="pb-2"> Name</span>
                        <input className="bg-white bg-opacity-10 p-2 rounded" value={name} required
                            onChange={handleNameChange} />
                        <div className="flex py-2 gap-4 bg-">
                            <Button onClick={toggleModal} text="Close" color="secondary" />
                            <Button text="Join" color="primary" onClick={handleJoinClick} />

                        </div>
                    </form>
                </div>
            </div>
        );
    } else {
        return null;
    }
};

export default Modal;