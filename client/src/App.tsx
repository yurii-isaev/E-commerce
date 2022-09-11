import React from 'react';
import { Container } from 'react-bootstrap';
import { Route, Routes } from 'react-router';
import './App.css';
import Footer from './components/footer';
import Header from './components/header';
import HomePage from './pages/home.page'
import LoginPage from './pages/login.page';
import RegistrationPage from './pages/registration.page';

// <div className="App">
function App() {
  return (
     <>
       <Header/>
       <main>
         <Container>
           <Routes>
             <Route path="/" element={<HomePage/>}/>
             <Route path="/login" element={<LoginPage/>}/>
             <Route path="/registration" element={<RegistrationPage/>}/>
           </Routes>
         </Container>
       </main>
       <Footer/>
     </>
  );
}

export default App
