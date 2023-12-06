'use client';
import { useState, useEffect } from 'react'
import Navbar from "../components/navbar";
import { Movie } from "../types";
import { moviesMock } from '../mock';
import { motion, useAnimate, AnimatePresence} from "framer-motion";


function Swipe() {

    const [movie, setMovie] = useState<Movie>(moviesMock[0]);
    const [dislike_movie, setDislikeMovie] = useState<Movie>();
    const [like_movie, setLikeMovie] = useState<Movie>();
    const [scope, animate] = useAnimate();
    const userId = localStorage.getItem('userId');

    const fetchRandomMovie = async (usedIndexes) => {
        try {
            const response = await fetch('http://localhost:5000/movies');
            const data = await response.json();

            // Create an array of all indexes
            const allIndexes = Array.from({ length: data.data.data.length }, (_, index) => index);

            // Exclude the usedIndexes from allIndexes
            const availableIndexes = allIndexes.filter(index => !usedIndexes.includes(index));

            // If there are available indexes, select a random one
            if (availableIndexes.length > 0) {
                const randomIndex = availableIndexes[Math.floor(Math.random() * availableIndexes.length)];
                const randomMovie = data.data.data[randomIndex];
                return { randomMovie, randomIndex };
            } else {
                // All indexes have been used, handle this case based on your requirements
                console.warn('No available indexes. All movies have been used.');
                return null;
            }
        } catch (error) {
            console.error('Error fetching movies data:', error);
            return null;
        }
    };



    const fetchData = async () => {
        try {
            const { randomMovie: movie, randomIndex: index1 } = await fetchRandomMovie([]);
            setMovie(movie);
            animate("img", {x:0});
            setTimeout(() => {
                animate("img", {opacity: 300});
            },400);
            
            
        } catch (error) {
            console.error('Error fetching movies data:', error);
        }
    };
    useEffect(() => {
        fetchData();
    }, []);

    const handleLike = async () => {
        try {

            
            const movieId = movie.id; 
            try {

                const response = await fetch(`http://localhost:5000/addliked/${userId}/${movieId}`, {
                    method: 'PATCH',
                    headers: {
                        'Content-Type': 'application/json',
                    }
                });

                const data = await response.json();
                console.log(data);
            } catch (error) {
                console.error('Error fetching adding liked movie', error);
            }
            fetchData();
        } catch (error) {
            console.error('Error liking movie:', error);
        }
    };

    function dislike(){
        animate("img", {x: -600});
        animate("img", {opacity: 0});
        
        
        setTimeout(fetchData, 300);
        
        setDislikeMovie(movie);
    }

    function like(){
        animate("img", {x: 600});
        animate("img", {opacity: 0});
        
        
        setTimeout(handleLike, 300);
       
        setLikeMovie(movie);
        
    }

    function LikeMovie(){
        if(like_movie){
            return <img className="border-white border-2  object-fill" src={like_movie.bgImg} />
        }
        return (
            <svg width="848" height="477">
                <rect className='object-fill bg-black' />
            </svg>
        )
    }

    function DislikeMovie(){
        if(dislike_movie){
            return <img className="border-white border-2 object-fill" src={dislike_movie.bgImg} />
        }
        return (
            <svg width="848" height="477">
                <rect className='object-fill bg-black' />
            </svg>
        )
    }

    function MovieGenres(){
        let i = 0;
        while(i < movie.genres.length){

        }
        
    }


    return (
        <section className="flex flex-col h-screen overflow-hidden">
            <Navbar />
            <div className="flex h-[100%]">
                <div className="skew-y-3 h-96 w-[1/3] flex-col flex gap-4 justify-end items-center rounded-xl overflow-hidden p-4 brightness-50 max-md:hidden">
                    Disliked Movies
                    <DislikeMovie />
                </div>
                
                
                <motion.div ref={scope} className="overflow-hidden relative flex flex-col gap-1 bg-black rounded-md  w-full h-full aspect-video border border-neutral-900 p-4">
                  
                    
                    <p className="text-center font-bold text-4xl pt-4">{movie.name}</p>
                    <img className="object-contain h-[20rem] border-white border-2 mt-4" src={movie.bgImg} alt={`${movie.name} background`} />
                    
                    <div className='flex flex-row gap-2 pt-4 '>
                    {movie.genres.map( (genres,index)=>
                        (
                            <p className=' border-2 border-white rounded-lg pl-2 pr-2'key={index}>{genres}</p>
                        )
                    )}
                    </div>
                    <p className=" w-full pt-4"> {movie.description}</p>
                    <div className='flex flex-row  gap-4   place-content-center w-full h-full pt-[3%]'>
                        <svg onClick={dislike} className='cursor-pointer' width="103" height="102" viewBox="0 0 63 62" fill="none" xmlns="http://www.w3.org/2000/svg">
                            <rect x="0.816895" width="62" height="62" rx="31" fill="white" fill-opacity="0.1"/>
                            <path d="M30.5119 39.1676C30.5931 40.3376 31.7019 41.2413 32.8969 40.9426L33.2231 40.8613C33.5224 40.7914 33.7976 40.6427 34.02 40.4306C34.2425 40.2185 34.4041 39.9507 34.4881 39.6551C34.7631 38.6351 35.1544 36.5151 34.5656 34.0176C34.7356 34.0426 34.9219 34.0638 35.1194 34.0813C36.0106 34.1626 37.2056 34.1701 38.2644 33.8176C38.9119 33.6013 39.5069 32.9676 39.7644 32.2276C39.8933 31.8694 39.934 31.4854 39.8831 31.1081C39.8322 30.7308 39.6912 30.3713 39.4719 30.0601C39.5444 29.9126 39.6006 29.7576 39.6444 29.6076C39.7406 29.2701 39.7856 28.8976 39.7856 28.5376C39.7856 28.1751 39.7406 27.8051 39.6444 27.4663C39.5953 27.292 39.5283 27.1231 39.4444 26.9626C39.6556 26.4788 39.5781 25.9376 39.4406 25.5263C39.2998 25.1207 39.0937 24.7407 38.8306 24.4013C38.8981 24.2101 38.9256 24.0101 38.9256 23.8201C38.9203 23.4192 38.8113 23.0265 38.6094 22.6801C38.1919 21.9463 37.3631 21.3501 36.1919 21.3501H31.8169C31.0606 21.3501 30.4794 21.4501 29.9844 21.6213C29.5587 21.7772 29.152 21.9806 28.7719 22.2276L28.7119 22.2638C28.0819 22.6488 27.4631 23.0263 26.1269 23.1676C25.1694 23.2688 24.3169 24.0426 24.3169 25.0988V30.0988C24.3169 31.1626 25.1731 31.8901 26.0131 32.1188C27.0744 32.4088 27.9806 33.1026 28.6781 33.8813C29.3781 34.6638 29.8206 35.4813 29.9769 35.9288C30.2256 36.6476 30.4219 37.8538 30.5119 39.1676V39.1676Z" fill="#EF4444"/>
                        </svg>
                        <div className='flex flex-row place-content-center pt-[50px] overflow-hidden'>
                        <svg width="277" height="12" viewBox="0 0 277 12" fill="none" xmlns="http://www.w3.org/2000/svg">
                            <path d="M0.816895 6L10.8169 11.7735V0.226497L0.816895 6ZM276.927 6L266.927 0.226497V11.7735L276.927 6ZM9.81689 7H267.927V5H9.81689V7Z" fill="white" fill-opacity="0.4"/>
                        </svg>

                        </div>
                        
                        <svg onClick={like} className='cursor-pointer' width="103" height="102" viewBox="0 0 63 62" fill="none" xmlns="http://www.w3.org/2000/svg">
                            <rect x="0.926758" width="62" height="62" rx="31" fill="white" fill-opacity="0.1"/>
                            <path d="M30.6218 23.1813C30.703 22.0126 31.8118 21.1088 33.0068 21.4063L33.333 21.4888C33.9118 21.6338 34.4255 22.0588 34.598 22.6951C34.873 23.7151 35.2643 25.8338 34.6755 28.3326C34.8596 28.3071 35.0442 28.2859 35.2293 28.2688C36.1205 28.1876 37.3155 28.1788 38.3743 28.5313C39.0218 28.7476 39.6168 29.3826 39.8743 30.1226C40.1043 30.7876 40.0743 31.5751 39.5818 32.2888C39.6543 32.4376 39.7105 32.5913 39.7543 32.7426C39.8505 33.0801 39.8955 33.4513 39.8955 33.8126C39.8955 34.1738 39.8505 34.545 39.7543 34.8826C39.7055 35.0513 39.6418 35.2238 39.5543 35.3876C39.7655 35.8713 39.688 36.4113 39.5505 36.8226C39.4097 37.2286 39.2037 37.609 38.9405 37.9488C39.008 38.1388 39.0355 38.3388 39.0355 38.5301C39.0355 38.9113 38.9243 39.3113 38.7193 39.67C38.3018 40.4025 37.473 41.0001 36.3018 41.0001H31.9268C31.1705 41.0001 30.5893 40.8988 30.0943 40.7275C29.6686 40.5721 29.2619 40.3692 28.8818 40.1226L28.8218 40.085C28.1918 39.7013 27.573 39.3238 26.2368 39.1825C25.2793 39.08 24.4268 38.3076 24.4268 37.2501V32.2501C24.4268 31.1876 25.283 30.4601 26.123 30.2313C27.1843 29.9413 28.0905 29.2476 28.788 28.4688C29.488 27.6851 29.9305 26.8688 30.0868 26.4201C30.3355 25.7013 30.5318 24.4963 30.6218 23.1826V23.1813Z" fill="#22C55E"/>
                        </svg>

                    </div>
                </motion.div>
                <div className=" -skew-y-3 h-96 w-[1/3] flex-col flex gap-4 justify-end items-center rounded-xl overflow-hidden p-4 brightness-50 max-md:hidden">
                    Liked Movies
                    <LikeMovie/>
                </div>
            </div>
            
        </section>
    );
};

export default Swipe
