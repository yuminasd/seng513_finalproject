import React, { useState, useEffect } from 'react';

const MovieCard = ({ movie }) => {
    const [isModalOpen, setIsModalOpen] = useState(false);

    const handleOpenModal = () => {
        setIsModalOpen(true);
    };

    const handleCloseModal = () => {
        setIsModalOpen(false);
    };

    // Add an event listener to close the modal when clicking outside
    useEffect(() => {
        const handleOutsideClick = (event) => {
            if (isModalOpen && event.target === event.currentTarget) {
                handleCloseModal();
            }
        };

        if (isModalOpen) {
            document.addEventListener('click', handleOutsideClick);
        }

        return () => {
            document.removeEventListener('click', handleOutsideClick);
        };
    }, [isModalOpen]);

    return (
        <div className="overflow-hidden relative flex flex-col bg-white bg-opacity-10 rounded-md h-[10rem]" onClick={handleOpenModal}>

            <img className="object-contain h-[10rem] " src={movie.img} alt={`${movie.name} background`} />

            {/* Modal */}
            {isModalOpen && (
                <div className="fixed top-0 left-0 w-full h-full bg-black bg-opacity-75 flex items-center justify-center z-50">
                    <div className="bg-white bg-opacity-10  p-4 rounded-md">
                        <p className="text-center">{movie.name}</p>
                        <button onClick={handleCloseModal}>Close</button>
                    </div>
                </div>
            )}
        </div>
    );
};

export default MovieCard;
