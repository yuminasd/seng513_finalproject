
import React, { useState } from 'react';
import Button from '../button';

interface ModalProps {
    hidden: boolean;
    onClose: () => void;
}

const Modal: React.FC<ModalProps> = ({ hidden, onClose }) => {
    const toggleModal = () => {
        onClose();
    };

    let code = "";

    if (!hidden) {
        return (
            <div className="fixed top-0 left-0 w-full h-full">
                <div className="absolute w-full h-full bg-black bg-opacity-50 flex justify-center items-center " >
                    <form className="bg-neutral-900 p-4 text-white flex flex-col z-50 rounded-xl">

                        <span className="pb-2"> Code</span>
                        <input className="bg-white bg-opacity-10 p-2 rounded" defaultValue={code} />
                        <div className="flex py-2 gap-4 bg-">
                            <Button onClick={toggleModal} text="Close" color="secondary" />
                            <Button text="Join" color="primary" />

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