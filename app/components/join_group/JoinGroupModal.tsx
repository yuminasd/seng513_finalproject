import React, { useState } from 'react';
import Modal from './modal';
import Button from '../button';


const JoinGroupModal: React.FC = () => {
    const [isModalOpen, setIsModalOpen] = useState(false);

    const openModal = () => {
        setIsModalOpen(true);
    };

    const closeModal = () => {
        setIsModalOpen(false);
    };

    return (
        <div>
            <Button onClick={openModal} text="Join Group" color="secondary" />
            <Modal hidden={!isModalOpen} onClose={closeModal} />
        </div>
    );
};

export default JoinGroupModal;