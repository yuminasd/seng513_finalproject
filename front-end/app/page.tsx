'use client';
import { useState, useEffect } from 'react'
import Navbar from "./components/navbar";
import JoinGroupModal from "./components/join_group/JoinGroupModal";
import Table from "./components/table";
import CreateGroupModal from "./components/create_group/CreateGroupModal";
import { Group, User } from './types';


function Home() {
  const [groups, setGroups] = useState<Group[]>([]);
  const [user, setUser] = useState<User | null>(null);
  const userId = typeof window !== 'undefined' ? localStorage.getItem('userId') : null;


  // Table Columns
  const columns = ['name', 'id'];

  useEffect(() => {

    const fetchData = async () => {
      try {
        console.log(userId);
        const response = await fetch('http://localhost:5000/users/' + userId);
        const data = await response.json();
        setGroups(data.data.data.groupID);
        setUser(data.data.data);
      } catch (error) {
        console.error('Error fetching groups data:', error);
      }
    };

    fetchData();
  }, []);

  return (
    <section >
      <Navbar />
      {/* <MockRead /> */}
      <div className="w-full  flex gap-4 p-4">
        <JoinGroupModal user={user} />
        <CreateGroupModal user={user} />
      </div>
      {groups ? (
        <Table columns={columns} data={groups} page="groups" />
      ) : (
        <p>Loading group data...</p>
      )}


    </section>
  );
};


export default Home;