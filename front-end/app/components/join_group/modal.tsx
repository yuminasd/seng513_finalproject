
import React, { useState } from 'react';
import Button from '../button';
import { User } from '@/app/types';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
interface ModalProps {
    hidden: boolean;
    onClose: () => void;
    user: User;
}

const Modal: React.FC<ModalProps> = ({ hidden, onClose, user }) => {


    const [code, setCode] = useState('');
    const handleCodeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setCode(event.target.value);
    };
    const toggleModal = () => {
        onClose();
    };

    const handleJoinClick = async () => {
        try {
            // event.preventDefault();
            console.log(user);
            const response = await fetch('http://localhost:5000/groups/' + code + '/members', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify([user.id]),
            });

            if (response.ok) {
                toast.success("Successfully Joined Group");
                setTimeout(() => {
                    // Handle successful response, e.g., close the modal
                    onClose();
                }, 2000);
            } else {
                // Handle error response
                toast.error('Failed to join group: ' + response.status + response.statusText);
            }
        } catch (error) {

            toast.error('Error joining group: ' + error);
        }
    };



    if (!hidden) {
        return (
            <div className="fixed top-0 left-0 w-full h-full">
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
                <div className="absolute w-full h-full bg-black bg-opacity-50 flex justify-center items-center " >
                    <form className="bg-neutral-900 p-4 text-white flex flex-col z-50 rounded-xl">

                        <span className="pb-2"> Code</span>
                        <input className="bg-white bg-opacity-10 p-2 rounded" value={code} required
                            onChange={handleCodeChange} />
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