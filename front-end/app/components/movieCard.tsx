import React, { useState } from 'react';
import { Movie } from '../types';

const MovieCard: React.FC<{ movie: Movie }> = ({ movie }) => {
    const [openModal, setOpenModal] = useState(false);

    const toggleModal = () => {
        setOpenModal(!openModal);
    }


    return (
        <div className="overflow-hidden relative flex flex-col bg-black rounded-md  w-full aspect-video" onClick={toggleModal}>

            <img className="object-contain hover:opacity-50" src={movie.img} alt={`${movie.name} background`} />

            {/* Modal */}
            {openModal ? (
                <div className="fixed top-0 left-0 w-full h-full bg-black bg-opacity-75 flex items-center justify-center z-50" onClick={toggleModal} >
                    <div className="bg-neutral-900 p-4 rounded-md flex flex-col gap-4 w-[40rem] rounded-xl" >

                        <p className="text-center">{movie.name}</p>
                        <img className="object-contain h-[20rem] " src={movie.img} alt={`${movie.name} background`} />
                        <p>Rating: {movie.rating}</p>
                        <p className="w-full"> This is a mock details explaining the cool mov fdafasdfasdfasdfasdfdsfasdfsddfsdfasdfasdfdsfadsfasdfdsf dfdf dfsa safasdfasdfie</p>
                    </div>
                </div>
            ) : null}
        </div>
    );
};

export default MovieCard;
