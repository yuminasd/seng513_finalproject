
import React, { useState } from 'react';
import Button from '../button';
import { User } from '@/app/types';
import MultiSelect from '../multiSelect';

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
    const [selectedGenres, setSelectedGenres] = useState([]);
    const handleGenreChange = (selectedOptions: React.SetStateAction<never[]>) => {
        setSelectedGenres(selectedOptions);
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
                    Genre: ['Canadian', 'Comedy Movies', 'TV Comedies', 'Mockumentaries', 'Late Night Comedy Movies', 'Stand-Up Comedy', 'British', "Kids' TV", 'Kids & Family Movies', 'Kids Music', 'TV Cartoons', 'Kids Movies', 'Music', 'Documentary Films', 'Concerts', 'Mystery Movies', 'Action & Adventure Movies', 'Drama Movies', 'Social Issue Dramas', 'Tearjerker Movies', 'Australian', 'Thriller Movies', 'Movies Based on Books', 'Independent Movies', 'Romantic Comedy Movies', 'Romantic Movies', 'Teen Movies', 'Family Movies', 'Horror Movies', 'TV Shows Based on Books', 'LGBTQ+ Movies', 'Musicals', 'Danish', 'Classic Movies', 'Satires', 'Period Pieces', 'Movies Based on Real Life', 'Courtroom Movies', 'Peruvian', 'Social & Cultural Docs', 'Irish', 'Sports Movies', 'Historical Documentaries', 'Sports & Fitness', 'Basketball Movies', 'Latin Music', 'Biographical Documentaries', 'German', 'Dutch', 'Nature & Ecology Documentaries', 'Science & Nature Docs', 'Spy Movies', 'Polish', 'Sci-Fi Movies', 'Belgian', 'Japanese', 'Family Watch Together TV', 'Italian', 'Brazilian', 'Faith & Spirituality', 'Spanish', 'Political Documentaries', 'LGBTQ+ Documentaries', 'Hip-Hop', 'True Crime Documentaries', 'Sports Dramas', 'Film Noir', 'Colombian', 'Singaporean', 'French', 'Adult Animation', 'Indian', 'Nollywood', 'African Movies', 'Sci-Fi Anime', 'Anime Movies', 'Swedish', 'Quirky Romance', 'Mexican', 'Vampire Horror Movies', 'Filipino', 'Chilean', 'Thai', 'South African', 'Variety TV', 'Anthology Films', 'Hindi-Language Movies', 'Bollywood Movies', 'Malaysian', 'Monster Movies', 'Turkish', 'Egyptian', 'Middle Eastern Movies'],
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
                    <form className="bg-neutral-900 p-4 text-white flex flex-col gap-2 z-50 rounded-xl">

                        <span className="pb-2"> Name</span>
                        <input className="bg-white bg-opacity-10 p-2 rounded" value={name} required
                            onChange={handleNameChange} />

                        <MultiSelect selectedGenres={selectedGenres} onChange={handleGenreChange} darkMode={true} />

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