'use client';
import { useState, useEffect } from 'react'
import Button from "./components/button";
import Navbar from "./components/navbar";
import JoinGroupModal from "./components/join_group/JoinGroupModal";
import Table from "./components/table";
import CreateGroupModal from "./components/create_group/CreateGroupModal";
import { groupsMock } from "./mock";
import { MockRead } from "./functions/users/read";
import { Group, User } from './types';


function Home() {
  const [groups, setGroups] = useState<Group[]>([]);
  const [user, setUser] = useState<User | null>(null);
  // Table Columns
  const columns = ['name', 'id'];

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:5000/users/656500413c49f1af1a59b5d1');
        const data = await response.json();
        // console.log(data.data.data);
        setGroups(data.data.data.groupID);
        setUser(data.data.data);
      } catch (error) {
        console.error('Error fetching groups data:', error);
      }
    };

    fetchData();
  }, []);

  return (
    <section className="h-screen overflow-y-auto pb-36">
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