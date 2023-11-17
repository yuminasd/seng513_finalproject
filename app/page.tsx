'use client';
import Button from "./components/button";
import Navbar from "./components/navbar";
import JoinGroupModal from "./components/join_group/JoinGroupModal";
import Table from "./components/table";
import CreateGroupModal from "./components/create_group/CreateGroupModal";

interface Group {
  GroupName: string;
  Code: string;
  users: User[];
  movies: [];
}

interface User {
  name: string;
  img: string;
}

function Home() {
  // Table Columns
  const columns = ['GroupName', 'Code'];
  //Mock Data
  const groups: Group[] = [
    {
      GroupName: "Group1",
      Code: "123",
      users: [
        {
          name: "person1",
          img: "",
        }
      ],
      movies: [],
    },
    {
      GroupName: "Group2",
      Code: "456",
      users: [],
      movies: [],
    },
  ];

  //Mock Button Click
  const handleButtonClick = () => {
    alert('Button clicked!');
  };

  return (
    <main className="flex min-h-screen flex-col  pt-16">
      <Navbar />

      {/* {groups.map((group, index) => (
        // <a key={index} href={`/group/${group.groupCode}`}>
        <a key={index} href={`/group`}>
          {group.groupName}
          {group.groupCode}
        </a>
      ))} */}



      <div className="w-full  flex gap-4 p-4">
        <JoinGroupModal />
        <CreateGroupModal />
      </div>

      <Table columns={columns} data={groups} />
    </main>
  );
};


export default Home;