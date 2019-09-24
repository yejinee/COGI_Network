const bcrypt = require('bcrypt');
const Sequelize = require('sequelize');
// const jwtSecret =  require('./jwtConfig');
// const BCRPYT_SALT_ROUNDS = 12;

const Op = Sequelize.Op;

const passport = require('passport');
const LocalStrategy = require('passport-local').Strategy;
// const JWTstrategy = require('passport-jwt').Strategy;
// const ExtractJWT = require('passport-jwt').ExtractJwt;
const User = require('../models/User');


passport.use(
    'signup',
    new LocalStrategy(
        {
            usernameField : 'email',
            passwordField: 'PASSWORD',
            passReqToCallback: true,
            session: false,
        },
        (req, username, password, done) => {
            try {
                User.findOne({
                    where: {
                        [Op.or]: [
                            {
                                email: username,
                            }
                        ],
                    },
                }).then(user => {
                    if (user != null) {
                        console.log('username or email already taken');
                        return done(null, false, {
                            message: 'username or email already taken',
                        });
                    }
                    return done(null, user);
                });
            } catch(err) {
                return done(err);
            }
        },
    ),
);


passport.use(
    'login',
    new LocalStrategy(
        {
            usernameField : 'email',
            passwordField : 'PASSWORD',
            session: false,
        },
        (username, password, done) => {
            try{
                User.findOne({
                    where: {
                        email : username,
                    },
                }).then(user => {
                    if(user === null){
                        return done(null, false, { message: 'bad username'});
                    }
                    bcrypt.compare(password, user.PASSWORD).then(response => {
                        if (response !== true) {
                            console.log('passwords do not match');
                            return done(null, false, {message: 'passwords do not match'});
                        }
                        console.log('user found & authenticated');
                        return done(null, user);
                    });
                });
            } catch (err){
                done(err);
            }
        },
    ),
);