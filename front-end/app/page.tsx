'use client';
import Button from "./components/button";
import Navbar from "./components/navbar";
import JoinGroupModal from "./components/join_group/JoinGroupModal";
import Table from "./components/table";
import CreateGroupModal from "./components/create_group/CreateGroupModal";
import { groupsMock } from "./mock";
import { MockRead } from "./functions/users/read";


function Home() {

  // Table Columns
  const columns = ['groupName', 'code'];


  //Mock Button Click
  const handleButtonClick = () => {
    alert('Button clicked!');
  };

  return (
    <section >
      <Navbar />
      <MockRead />
      <div className="w-full  flex gap-4 p-4">
        <JoinGroupModal />
        <CreateGroupModal />
      </div>

      <Table columns={columns} data={groupsMock} page="groups" />
    </section>
  );
};


export default Home;