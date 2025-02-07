import React, { useEffect, useState } from 'react';

interface HomeProps {}

const Home: React.FC<HomeProps> = () => {
  const [message, setMessage] = useState('');
  useEffect(() => {
    fetch('http://localhost:8080/ping')
      .then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then((data) => {
        setMessage(data.message);
      })
      .catch((error) => {
        console.error('There was a problem with the fetch operation:', error);
        // Optionally, you can update state here to show an error message to the user
        setMessage('Error fetching data');
      });
  }, []);

  return (
    <div className="bg-blue-500 h-screen w-screen flex flex-col items-center justify-center">
      <header className=" h-10 w-full flex justify-center">
        <p className="text-xl font-bold text-white uppercase">
          {message || 'Waiting for response...'}
        </p>
      </header>
    </div>
  );
};

export default Home;
