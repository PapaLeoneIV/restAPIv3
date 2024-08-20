import { Inter } from "next/font/google";
import ButtonList from "./components/ButtonList";
import Header from "./components/Header";
import PostForm from "./components/Forms/PostForm";
import Logo from "./components/Logo42";
import {ShowSingle} from "./components/Forms/ShowSingle";
import { useState } from "react";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {

  const [currentForm, setCurrentForm] = useState('general');

  const handleFormChange = (form : string)=>{
    setCurrentForm(form);
  }

  return (
    <div>
     <Header/>
      <main className="flex flex-row h-screen w-screen bg-black text-white pt-20 items-center">
        <div className="w-1/12"></div>
        <div className="flex flex-column">
          <ButtonList onFormChange={handleFormChange} />
        {(currentForm === 'post' || currentForm === "general") && <PostForm />}
        {(currentForm === 'single') && <ShowSingle/>}
        {(currentForm === 'all') && <PostForm />}
        {(currentForm === 'modify') && <PostForm />}
        {(currentForm === 'update') && <PostForm />}
        </div>
        <Logo />
      </main>
    </div>
  );
}
