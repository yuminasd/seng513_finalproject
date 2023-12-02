'use client';
import { useState, useEffect } from 'react'
import Button from "./components/button";
import Navbar from "./components/navbar";
import JoinGroupModal from "./components/join_group/JoinGroupModal";
import Table from "./components/table";
import CreateGroupModal from "./components/create_group/CreateGroupModal";
import { groupsMock } from "./mock";
import { MockRead } from "./functions/users/read";

interface Group {
  groupName: string;
  code: string;
  // Add more properties if needed
}
function Home() {
  const [groups, setGroups] = useState<Group[]>([]);
  // Table Columns
  const columns = ['groupName', 'code'];

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:5000/groups/65692dca06a2d9ee2acd91e4');
        const data = await response.json();
        console.log(data.data.data);
        setGroups(data);
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
        <JoinGroupModal />
        <CreateGroupModal />
      </div>

      <Table columns={columns} data={groupsMock} page="groups" />
    </section>
  );
};


export default Home;