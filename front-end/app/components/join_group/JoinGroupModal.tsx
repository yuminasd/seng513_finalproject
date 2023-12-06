import React, { useState } from 'react';
import Modal from './modal';
import Button from '../button';
import { User } from '@/app/types';

interface CreateGroupModalProps {
    user: User | null;
}


const JoinGroupModal: React.FC<CreateGroupModalProps> = ({ user }) => {
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
            <Modal hidden={!isModalOpen} onClose={closeModal} user={user} />
        </div>
    );
};

export default JoinGroupModal;