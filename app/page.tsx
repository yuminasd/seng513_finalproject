
import Navbar from "./components/navbar";

interface Group {
  groupName: string;
  groupCode: string;
  users: User[];
  movies: [];
}

interface User {
  name: string;
  img: string;
}

function Home() {
  const groups: Group[] = [
    {
      groupName: "Group1",
      groupCode: "123",
      users: [
        {
          name: "person1",
          img: "",
        }
      ],
      movies: [],
    },
    {
      groupName: "Group2",
      groupCode: "456",
      users: [],
      movies: [],
    },
    // Add more groups as needed
  ];

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <Navbar />
      <h1>Groups page</h1>
      {groups.map((group, index) => (
        // <a key={index} href={`/group/${group.groupCode}`}>
        <a key={index} href={`/group`}>
          {group.groupName}
          {group.groupCode}
        </a>
      ))}
    </main>
  );
};


export default Home;