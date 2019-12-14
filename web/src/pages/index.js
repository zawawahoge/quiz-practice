import React from 'react'
import Link from 'next/link'
import Head from 'next/head'
import Nav from '../components/nav'

const Home = () => (
  <div>
    <Head>
      <title>Home</title>
    </Head>

    <div>
      <ul>
        <li>
          <Link href='/word-list'><a>練習モード</a></Link>
        </li>
      </ul>
    </div>

  </div>
)

export default Home
