import { Inter } from "next/font/google";
import ButtonList from "./components/ButtonList";
import Header from "./components/Header";
import ComplaintForm from "./components/ComplaintForm";
import Logo from "./components/Logo42";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    <div>
     <Header/>
      <main className="flex flex-row h-screen w-screen bg-black text-white pt-20 items-center">
        <div className="w-1/12"></div>
        <ButtonList/>
        <ComplaintForm />
        <Logo />
      </main>
    </div>
  );
}
